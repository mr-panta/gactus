package gactus

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/rpc"

	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/gactus/service"
)

type GactusService interface {
	Start() error
	RegisterProcessors(processors []*Processor)
	SendRequest(ctx context.Context, command string, req proto.Message, res proto.Message) error
}

type gactusService struct {
	tcpPort int
	service *service.Service
}

func NewGactusService(tcpPort int) GactusService {
	return &gactusService{
		tcpPort: tcpPort,
		service: service.NewService(),
	}
}

func (gs *gactusService) Start() error {
	lis, err := net.Listen(networkTCP, fmt.Sprintf(":%d", gs.tcpPort))
	if err != nil {
		return err
	}
	server := rpc.NewServer()
	err = server.Register(gs.service)
	if err != nil {
		return err
	}
	server.Accept(lis)
	return errors.New("service stopped")
}

func (gs *gactusService) RegisterProcessors(processors []*Processor) {
	for _, p := range processors {
		gs.service.AddProcessor(p.Command, &service.Processor{
			Req:            p.Req,
			Res:            p.Res,
			HTTPConfig:     p.HTTPConfig,
			HTTPMiddleware: p.HTTPMiddleware,
			Process:        p.Process,
		})
	}
}

func (gs *gactusService) SendRequest(ctx context.Context, command string, req proto.Message, res proto.Message) error {
	return gs.service.SendRequest(ctx, command, req, res)
}

type Processor struct {
	Command        string
	Req            proto.Message
	Res            proto.Message
	HTTPConfig     *pb.HttpConfig
	HTTPMiddleware func(ctx context.Context, header, query map[string]string, req, res proto.Message) error
	Process        func(ctx context.Context, req, res proto.Message) error
}
