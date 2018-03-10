package proxy

import (
	"fmt"
	"github.com/cjongseok/mtproto/proto"
	"github.com/cjongseok/slog"
	"google.golang.org/grpc"
	"net"
	"time"
)

type MProxy struct {
	server   *grpc.Server
	mmanager *mtp.MManager
	mconn    *mtp.MConn
	port     int
	streams  []UpdateStreamer_ListenOnUpdatesServer
}

func NewMProxy(port int) *MProxy {
	p := &MProxy{}
	server := grpc.NewServer([]grpc.ServerOption{}...)
	mtp.RegisterMtprotoServer(server, p)
	RegisterUpdateStreamerServer(server, p)
	p.server = server
	p.port = port
	return p
}

func (p *MProxy) SignIn() {
	// 1. key path as an argument
}

func (p *MProxy) Connect(config mtp.Configuration, phoneNumber, preferredAddr string) error { // open mrptoro
	var err error
	p.mmanager, err = mtp.NewManager(config)
	if err != nil {
		return fmt.Errorf("invalid configuration: %s", err)
	}
	p.mconn, err = p.mmanager.LoadAuthentication(phoneNumber, preferredAddr)
	if err != nil {
		return fmt.Errorf("load auth failure: %s", err)
	}
	p.mconn.AddUpdateCallback(p)
	return nil
}

func (p *MProxy) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", p.port))
	if err != nil {
		return fmt.Errorf("socket open failure:", err)
	}
	go func() {
		slog.Logln(p, "start serving gRPC")
		err := p.server.Serve(lis)
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
				startErr := p.Start()
				if startErr != nil {
					slog.Logln(p, "restart failure:", startErr)
				}
			}()
		}
	}()
	return nil
}

func (p *MProxy) Stop() {
	p.server.GracefulStop()
}

func (p *MProxy) LogPrefix() string {
	return "[proxy]"
}

func (p *MProxy) OnUpdate(mu mtp.MUpdate) {
	var pu *Update
	switch casted := mu.(type) {
	case *mtp.PredUpdateShortMessage:
		pu = &Update{&Update_UpdateShortMessage{casted}}
	case *mtp.PredUpdatesDifference:
		pu = &Update{&Update_UpdatesDifference{casted}}
	default:
		slog.Logln(p, "unknown update:", mu)
	}
	if pu != nil {
		for _, s := range p.streams {
			if err := s.Send(pu); err != nil {
				// TODO: monitor the channels condition and reset abnormal ones
				slog.Logln(p, "stream an update failure:", err)
			}
		}
	}
}

func (p *MProxy) ListenOnUpdates(req *ListenRequest, stream UpdateStreamer_ListenOnUpdatesServer) error {
	p.streams = append(p.streams, stream)
	return nil
}
