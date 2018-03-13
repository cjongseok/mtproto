package proxy

import (
	"fmt"
	"github.com/cjongseok/mtproto/core"
	"github.com/cjongseok/slog"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
	grpcServer *grpc.Server
	mmanager   *core.Manager
	mconn      *core.Conn
	port       int
	streams    []chan *Update
}

func NewServer(port int) *Server {
	p := &Server{}
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	core.RegisterMtprotoServer(grpcServer, &core.RPCaller{p})
	RegisterUpdateStreamerServer(grpcServer, p)
	p.grpcServer = grpcServer
	p.port = port
	return p
}

func (p *Server) InvokeBlocked(msg core.TL) (interface{}, error) {
	return p.mconn.InvokeBlocked(msg)
}

func (p *Server) Start(config core.Configuration, phoneNumber, preferredAddr string) error {
	err := p.connect(config, phoneNumber, preferredAddr)
	if err != nil {
		return err
	}
	err = p.serve()
	if err != nil {
		return err
	}
	return nil
}

func (p *Server) connect(config core.Configuration, phoneNumber, preferredAddr string) error { // open mrptoro
	var err error
	p.mmanager, err = core.NewManager(config)
	if err != nil {
		return fmt.Errorf("invalid configuration: %s", err)
	}
	p.mconn, err = p.mmanager.LoadAuthentication(phoneNumber, preferredAddr)
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
	go func() {
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
			go func() {
				startErr := p.serve()
				if startErr != nil {
					slog.Logln(p, "restart failure:", startErr)
				}
			}()
		}
	}()
	return nil
}

func (p *Server) Stop() {
	p.grpcServer.GracefulStop()
}

func (p *Server) LogPrefix() string {
	return "[proxy]"
}

func (p *Server) OnUpdate(mu core.Update) {
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
	core.MtprotoClient
	UpdateStreamerClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		return nil, err
	}
	corerotoClient := core.NewMtprotoClient(conn)
	updateClient := NewUpdateStreamerClient(conn)
	return &Client{corerotoClient, updateClient}, nil
}

func toProxyUpdate(u core.Update) *Update {
	switch pu := u.(type) {
	case *core.PredUpdatesState:
		return &Update{&Update_UpdatesState{pu}}
	case *core.PredUpdateShortMessage:
		return &Update{&Update_UpdateShortMessage{pu}}
	case *core.PredUpdateShortChatMessage:
		return &Update{&Update_UpdateShortChatMessage{pu}}
	case *core.PredUpdateShort:
		return &Update{&Update_UpdateShort{pu}}
	case *core.PredUpdates:
		return &Update{&Update_Updates{pu}}
	case *core.PredUpdateShortSentMessage:
		return &Update{&Update_UpdateShortSentMessage{pu}}
	case *core.PredUpdatesDifference:
		return &Update{&Update_UpdatesDifference{pu}}
	case *core.PredUpdatesDifferenceSlice:
		return &Update{&Update_UpdatesDifferenceSlice{pu}}
	case *core.PredUpdateNewMessage:
		return &Update{&Update_UpdateNewMessage{pu}}
	case *core.PredUpdateReadMessagesContents:
		return &Update{&Update_UpdateReadMessagesContents{pu}}
	case *core.PredUpdateDeleteMessages:
		return &Update{&Update_UpdateDeleteMessages{pu}}
	case *core.PredUpdateNewEncryptedMessage:
		return &Update{&Update_UpdateNewEncryptedMessage{pu}}
	case *core.PredUpdateChannel:
		return &Update{&Update_UpdateChannel{pu}}
	case *core.PredUpdateChannelMessageViews:
		return &Update{&Update_UpdateChannelMessageViews{pu}}
	case *core.PredUpdateChannelTooLong:
		return &Update{&Update_UpdateChannelTooLong{pu}}
	case *core.PredUpdateReadChannelInbox:
		return &Update{&Update_UpdateReadChannelInbox{pu}}
	case *core.PredUpdateReadChannelOutbox:
		return &Update{&Update_UpdateReadChannelOutbox{pu}}
	case *core.PredUpdateNewChannelMessage:
		return &Update{&Update_UpdateNewChannelMessage{pu}}
	}
	return nil
}
