package gactus

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mr-panta/gactus/internal/util"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus/internal/config"
	"github.com/mr-panta/gactus/internal/service"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
)

type gactusService struct {
	name            string
	coreAddr        string
	tcpPort         int
	handler         service.Handler
	minConns        int
	maxConns        int
	idleConnTimeout int
	waitConnTimeout int
	clearPeriod     int
}

// NewService [TOWRITE]
func NewService(name, coreAddr string, tcpPort, minConns, maxConns, idleConnTimeout, waitConnTimeout,
	clearPeriod int) (Service, error) {

	handler, err := service.NewHandler(coreAddr, minConns, maxConns, idleConnTimeout, waitConnTimeout, clearPeriod)
	if err != nil {
		return nil, err
	}
	return &gactusService{
		name:            name,
		coreAddr:        coreAddr,
		tcpPort:         tcpPort,
		handler:         handler,
		minConns:        minConns,
		maxConns:        maxConns,
		idleConnTimeout: idleConnTimeout,
		waitConnTimeout: waitConnTimeout,
		clearPeriod:     clearPeriod,
	}, nil
}

func (c *gactusService) listenTCP() error {
	ctx := context.Background()
	logger.Debugf(ctx, "start tcp server on port %d", c.tcpPort)
	return service.ListenAndServe(fmt.Sprintf(":%d", c.tcpPort), c.handler)
}

// Start [TOWRITE]
func (c *gactusService) Start() {
	ctx := context.Background()
	go func() {
		err := c.listenTCP()
		if err != nil {
			logger.Fatalf(ctx, err.Error())
		}
	}()
}

// Wait [TOWRITE]
func (c *gactusService) Wait() {
	p := make(chan os.Signal, 1)
	signal.Notify(p, os.Interrupt, syscall.SIGTERM)
	<-p
	logger.Warnf(context.Background(), "service server is terminated")
	os.Exit(0)
}

// RegisterService [TOWRITE]
func (c *gactusService) RegisterService(processors []*Processor) error {
	ctx := logger.GetContextWithLogID(context.Background(), c.name)
	logger.Debugf(ctx, "start registering service")
	addrs, err := util.GetIPAddrs()
	if err != nil {
		return err
	}
	for i := range addrs {
		addrs[i] = fmt.Sprintf("%s:%d", addrs[i], c.tcpPort)
	}
	req := &pb.RegisterServiceRequest{
		Addresses:           addrs,
		ProcessorRegistries: make([]*pb.ProcessorRegistry, len(processors)),
		ConnConfig: &pb.ConnectionConfig{
			MinConns:        uint32(c.minConns),
			MaxConns:        uint32(c.maxConns),
			IdleConnTimeout: uint32(c.idleConnTimeout),
			WaitConnTimeout: uint32(c.waitConnTimeout),
			ClearPeriod:     uint32(c.clearPeriod),
		},
	}
	for i, processor := range processors {
		req.ProcessorRegistries[i] = &pb.ProcessorRegistry{
			Command:    processor.Command,
			HttpConfig: processor.HTTPConfig,
		}
		c.handler.SetProcessor(processor.Command, &service.Processor{
			Req:            processor.Req,
			Res:            processor.Res,
			HTTPMiddleware: processor.HTTPMiddleware,
			Process:        processor.Process,
		})
	}
	res := &pb.RegisterServiceResponse{}
	code := c.SendRequest(ctx, config.CMDCoreRegisterService, req, res)
	if code != uint32(pb.Constant_RESPONSE_OK) {
		return errors.New(res.GetDebugMessage())
	}
	return nil
}

// SendRequest [TOWRITE]
func (c *gactusService) SendRequest(ctx context.Context, command string, req, res proto.Message) (code uint32) {
	var err error
	if logger.GetLogID(ctx) == "" {
		ctx = logger.GetContextWithLogID(ctx, command)
	}
	code, err = c.handler.SendRequest(logger.GetLogID(ctx), command, req, res)
	if err != nil {
		logger.Errorf(ctx, err.Error())
	}
	return code
}
