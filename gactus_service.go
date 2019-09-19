package gactus

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus/internal/config"
	"github.com/mr-panta/gactus/internal/service"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
)

type gactusService struct {
	name            string
	coreAddr        string
	tcpAddr         string
	minConns        int
	maxConns        int
	idleConnTimeout int
	waitConnTimeout int
	clearPeriod     int
	handler         service.Handler
}

// NewService [TOWRITE]
func NewService(name, coreAddr, tcpAddr string, minConns, maxConns, idleConnTimeout, waitConnTimeout,
	clearPeriod int) (Service, error) {

	handler, err := service.NewHandler(coreAddr, 0, 1, idleConnTimeout, waitConnTimeout, clearPeriod)
	if err != nil {
		return nil, err
	}
	return &gactusService{
		name:            name,
		coreAddr:        coreAddr,
		tcpAddr:         tcpAddr,
		minConns:        minConns,
		maxConns:        maxConns,
		idleConnTimeout: idleConnTimeout,
		waitConnTimeout: waitConnTimeout,
		clearPeriod:     clearPeriod,
		handler:         handler,
	}, nil
}

func (c *gactusService) listenTCP() error {
	ctx := context.Background()
	logger.Debugf(ctx, "start tcp server on %s", c.tcpAddr)
	return service.ListenAndServe(c.tcpAddr, c.handler)
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

// RegisterProcessors [TOWRITE]
func (c *gactusService) RegisterProcessors(processors []Processor) error {
	ctx := logger.GetContextWithLogID(context.Background(), c.name)
	logger.Debugf(ctx, "start registering processors")
	req := &pb.RegisterProcessorsRequest{
		Addr:                c.tcpAddr,
		ProcessorRegistries: make([]*pb.ProcessorRegistry, len(processors)),
	}
	for i, processor := range processors {
		req.ProcessorRegistries[i] = &pb.ProcessorRegistry{
			Command:    processor.GetCommand(),
			HttpConfig: processor.GetHTTPConfig(),
		}
		c.handler.SetProcess(processor.GetCommand(), processor.Process)
	}
	res := &pb.RegisterProcessorsResponse{}
	code := c.SendRequest(ctx, config.CMDCoreRegisterProcessors, req, res)
	if code != uint32(pb.Constant_RESPONSE_OK) {
		return errors.New(res.GetDebugMessage())
	}
	return nil
}

// SendRequest [TOWRITE]
func (c *gactusService) SendRequest(ctx context.Context, command string, req, res proto.Message) (code uint32) {
	var err error
	code, err = c.handler.SendCoreRequest(logger.GetLogID(ctx), command, req, res)
	if err != nil {
		logger.Errorf(ctx, err.Error())
	}
	return code
}
