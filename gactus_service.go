package gactus

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus/pkg/logger"
	"github.com/mr-panta/gactus/pkg/service"
	"github.com/mr-panta/gactus/pkg/tcp"
	pb "github.com/mr-panta/gactus/proto"
)

type defaultService struct {
	name             string
	coreAddr         string
	tcpAddr          string
	coreConnPoolSize int
	handler          service.Handler
}

func newDefaultService(name, coreAddr, tcpAddr string, coreConnPoolSize int) *defaultService {
	handler := service.NewHandler(coreAddr, coreConnPoolSize)
	return &defaultService{
		name:             name,
		coreAddr:         coreAddr,
		tcpAddr:          tcpAddr,
		coreConnPoolSize: coreConnPoolSize,
		handler:          handler,
	}
}

func (c *defaultService) listenTCP() error {
	ctx := context.Background()
	logger.Debugf(ctx, "start TCP server at %s", c.tcpAddr)
	return tcp.ListenAndServe(c.tcpAddr, c.handler)
}

// Start [TOWRITE]
func (c *defaultService) Start() {
	ctx := context.Background()
	go func() {
		err := c.listenTCP()
		if err != nil {
			logger.Fatalf(ctx, err.Error())
		}
	}()
}

// Wait [TOWRITE]
func (c *defaultService) Wait() {
	p := make(chan os.Signal, 1)
	signal.Notify(p, os.Interrupt, syscall.SIGTERM)
	<-p
	logger.Warnf(context.Background(), "service server is terminated")
	os.Exit(0)
}

// RegisterProcessors [TOWRITE]
func (c *defaultService) RegisterProcessors(processors []Processor) error {
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
	ctx := logger.GetContextWithLogID(context.Background(), c.name)
	logger.Debugf(ctx, "start registering processors")
	res := &pb.RegisterProcessorsResponse{}
	code := c.SendRequest(ctx, CMDRegisterProcessors, req, res)
	if code != uint32(pb.Constant_RESPONSE_OK) {
		return errors.New(res.GetDebugMessage())
	}
	return nil
}

// SendRequest [TOWRITE]
func (c *defaultService) SendRequest(ctx context.Context, command string, req, res proto.Message) (code uint32) {
	var err error
	code, err = c.handler.SendCoreRequest(logger.GetLogID(ctx), command, req, res)
	if err != nil {
		logger.Errorf(ctx, err.Error())
	}
	return code
}