package proxy

import (
	"fmt"
	"github.com/cjongseok/mtproto"
	"github.com/cjongseok/slog"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

type Server struct {
	grpcServer *grpc.Server
	mmanager   *mtproto.Manager
	mconn      *mtproto.Conn
	port       int
	streams    []chan *Update
	wg         *sync.WaitGroup
}

func NewServer(port int) *Server {
	p := &Server{}
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	mtproto.RegisterMtprotoServer(grpcServer, &mtproto.RPCaller{p})
	RegisterUpdateStreamerServer(grpcServer, p)
	p.grpcServer = grpcServer
	p.port = port
	p.wg = &sync.WaitGroup{}
	return p
}

func (p *Server) AddUpdateCallback(callback mtproto.UpdateCallback) {
	p.mconn.AddUpdateCallback(callback)
}

func (p *Server) InvokeBlocked(msg mtproto.TL) (interface{}, error) {
	return p.mconn.InvokeBlocked(msg)
}

func (p *Server) Start(config mtproto.Configuration) error {
	err := p.connect(config)
	if err != nil {
		return err
	}
	err = p.serve()
	if err != nil {
		return err
	}
	return nil
}

func (p *Server) connect(config mtproto.Configuration) error { // open mrptoro
	var err error
	p.mmanager, err = mtproto.NewManager(config)
	if err != nil {
		return fmt.Errorf("invalid configuration: %s", err)
	}
	p.mconn, err = p.mmanager.LoadAuthentication()
	if err != nil {
		return fmt.Errorf("load auth failure: %v", err)
	}
	p.mconn.AddUpdateCallback(p)
	return nil
}

func (p *Server) serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", p.port))
	if err != nil {
		return fmt.Errorf("socket open failure: %v", err)
	}
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		slog.Logln(p, "start serving gRPC")
		err := p.grpcServer.Serve(lis)
		switch err {
		case grpc.ErrServerStopped:
			slog.Logln(p, "gRPC stopped:", err)
		default:
			slog.Logln(p, "unknown error on serving gRPC:", err)
			slog.Logln(p, "cool down the proxy for a second...")
			<-time.After(time.Second)
			slog.Logln(p, "shut down the socket and restart the proxy...")
			lis.Close()
			p.wg.Add(1)
			go func() {
				p.wg.Done()
				startErr := p.serve()
				if startErr != nil {
					slog.Logln(p, "restart failure:", startErr)
				}
			}()
		}
	}()
	return nil
}

func (p *Server) Wait() {
	p.wg.Wait()
}

func (p *Server) Stop() {
	p.grpcServer.GracefulStop()
}

func (p *Server) LogPrefix() string {
	return "[proxy]"
}

func (p *Server) OnUpdate(mu mtproto.Update) {
	pu := toProxyUpdate(mu)
	if pu != nil {
		for _, s := range p.streams {
			go func() {
				s <- pu
			}()
		}
	}
}

func (p *Server) ListenOnUpdates(req *ListenRequest, stream UpdateStreamer_ListenOnUpdatesServer) error {
	ch := make(chan *Update)
	//TODO: thread safe streams
	p.streams = append(p.streams, ch)
	for {
		select {
		case u := <-ch:
			if err := stream.Send(u); err != nil {
				// TODO: monitor the channels condition and reset abnormal ones
				slog.Logln(p, "stream an update failure:", err)
			}
		}
	}
	return nil
}

type Client struct {
	mtproto.MtprotoClient
	UpdateStreamerClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		return nil, err
	}
	mtprotorotoClient := mtproto.NewMtprotoClient(conn)
	updateClient := NewUpdateStreamerClient(conn)
	return &Client{mtprotorotoClient, updateClient}, nil
}

func toProxyUpdate(u mtproto.Update) *Update {
	switch pu := u.(type) {
	case *mtproto.PredUpdatesState:
		return &Update{&Update_UpdatesState{pu}}
	case *mtproto.PredUpdateShortMessage:
		return &Update{&Update_UpdateShortMessage{pu}}
	case *mtproto.PredUpdateShortChatMessage:
		return &Update{&Update_UpdateShortChatMessage{pu}}
	case *mtproto.PredUpdateShort:
		return &Update{&Update_UpdateShort{pu}}
	case *mtproto.PredUpdates:
		return &Update{&Update_Updates{pu}}
	case *mtproto.PredUpdateShortSentMessage:
		return &Update{&Update_UpdateShortSentMessage{pu}}
	case *mtproto.PredUpdatesDifference:
		return &Update{&Update_UpdatesDifference{pu}}
	case *mtproto.PredUpdatesDifferenceSlice:
		return &Update{&Update_UpdatesDifferenceSlice{pu}}
	case *mtproto.PredUpdateNewMessage:
		return &Update{&Update_UpdateNewMessage{pu}}
	case *mtproto.PredUpdateReadMessagesContents:
		return &Update{&Update_UpdateReadMessagesContents{pu}}
	case *mtproto.PredUpdateDeleteMessages:
		return &Update{&Update_UpdateDeleteMessages{pu}}
	case *mtproto.PredUpdateNewEncryptedMessage:
		return &Update{&Update_UpdateNewEncryptedMessage{pu}}
	case *mtproto.PredUpdateChannel:
		return &Update{&Update_UpdateChannel{pu}}
	case *mtproto.PredUpdateChannelMessageViews:
		return &Update{&Update_UpdateChannelMessageViews{pu}}
	case *mtproto.PredUpdateChannelTooLong:
		return &Update{&Update_UpdateChannelTooLong{pu}}
	case *mtproto.PredUpdateReadChannelInbox:
		return &Update{&Update_UpdateReadChannelInbox{pu}}
	case *mtproto.PredUpdateReadChannelOutbox:
		return &Update{&Update_UpdateReadChannelOutbox{pu}}
	case *mtproto.PredUpdateNewChannelMessage:
		return &Update{&Update_UpdateNewChannelMessage{pu}}
	}
	return nil
}
