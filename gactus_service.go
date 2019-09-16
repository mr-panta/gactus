package gactus

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/mr-panta/gactus/pkg/logger"
	"github.com/mr-panta/gactus/pkg/service"
	"github.com/mr-panta/gactus/pkg/tcp"
)

type defaultService struct {
	coreAddr         string
	tcpAddr          string
	coreConnPoolSize int
	handler          *service.Handler
}

func newDefaultService(coreAddr, tcpAddr string, coreConnPoolSize int) *defaultService {
	return &defaultService{
		coreAddr:         coreAddr,
		tcpAddr:          tcpAddr,
		coreConnPoolSize: coreConnPoolSize,
	}
}

// Start [TOWRITE]
func (c *defaultService) Start() {
	ctx := context.Background()
	c.handler = service.NewHandler(c.coreAddr, c.coreConnPoolSize)
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
	// TODO: send processor data to core via TCP
	return nil
}

func (c *defaultService) listenTCP() error {
	ctx := context.Background()
	logger.Infof(ctx, "start TCP server at %s", c.tcpAddr)
	return tcp.ListenAndServe(c.tcpAddr, c.handler)
}
