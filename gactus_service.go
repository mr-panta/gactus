package gactus

import (
	"context"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"

	"github.com/mr-panta/gactus/config"
	"github.com/mr-panta/go-logger"

	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/gactus/service"
	"github.com/mr-panta/gactus/util"
)

type GactusService interface {
	Start()
	Wait()
	RegisterProcessors(processors []*Processor) error
	SendRequest(ctx context.Context, command string, req proto.Message, res proto.Message) error
}

type gactusService struct {
	name        string
	tcpPort     int
	coreAddress string
	service     *service.Service
}

type Processor struct {
	Command        string
	Req            proto.Message
	Res            proto.Message
	HTTPConfig     *pb.HttpConfig
	HTTPMiddleware func(ctx context.Context, header, query map[string]string, req, res proto.Message) error
	Process        func(ctx context.Context, req, res proto.Message) error
}

func NewGactusService(name string, coreAddress string, tcpPort int) GactusService {
	gs := &gactusService{
		name:        name,
		coreAddress: coreAddress,
		tcpPort:     tcpPort,
		service:     service.NewService(),
	}
	return gs
}

func (gs *gactusService) Start() {
	ctx := logger.GetContextWithLogID(context.Background(), "start_service")
	lis, err := net.Listen(networkTCP, fmt.Sprintf(":%d", gs.tcpPort))
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}
	server := rpc.NewServer()
	err = server.Register(gs.service)
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}
	err = gs.service.SetAddressCommands(gs.coreAddress, []string{
		config.CMDCoreRegisterService,
	})
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}
	go server.Accept(lis)
}

func (gs *gactusService) Wait() {
	p := make(chan os.Signal, 1)
	signal.Notify(p, os.Interrupt, syscall.SIGTERM)
	<-p
	logger.Warnf(context.Background(), "core server is terminated")
	os.Exit(0)
}

func (gs *gactusService) RegisterProcessors(processors []*Processor) error {
	ctx := logger.GetContextWithLogID(context.Background(), "register_processors")
	addrs, err := util.GetIPAddrs()
	if err != nil {
		return err
	}
	for i := range addrs {
		addrs[i] = fmt.Sprintf("%s:%d", addrs[i], gs.tcpPort)
	}
	processors = append(processors, &Processor{
		Command: config.CMDServiceUpdateRegistries,
		Req:     &pb.UpdateRegistriesRequest{},
		Res:     &pb.UpdateRegistriesResponse{},
		Process: gs.ProcessUpdateRegistries,
	})
	req := &pb.RegisterServiceRequest{
		Addresses:           addrs,
		ProcessorRegistries: make([]*pb.ProcessorRegistry, len(processors)),
	}
	for i, p := range processors {
		req.ProcessorRegistries[i] = &pb.ProcessorRegistry{
			Command:    p.Command,
			HttpConfig: p.HTTPConfig,
		}
		gs.service.AddProcessor(p.Command, &service.Processor{
			Req:            p.Req,
			Res:            p.Res,
			HTTPConfig:     p.HTTPConfig,
			HTTPMiddleware: p.HTTPMiddleware,
			Process:        p.Process,
		})
	}
	res := &pb.RegisterServiceResponse{}
	err = gs.service.SendRequest(ctx, config.CMDCoreRegisterService, req, res)
	if err != nil {
		return err
	}
	return nil
}

func (gs *gactusService) SendRequest(ctx context.Context, command string, req proto.Message, res proto.Message) error {
	return gs.service.SendRequest(ctx, command, req, res)
}
