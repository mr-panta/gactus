package gactus

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mr-panta/gactus/internal/core"
	"github.com/mr-panta/go-logger"
)

type gactusCore struct {
	httpAddr string
	tcpAddr  string
	handler  core.Handler
}

// NewCore [TOWRITE]
func NewCore(httpAddr, tcpAddr string) Core {
	return &gactusCore{
		httpAddr: httpAddr,
		tcpAddr:  tcpAddr,
		handler:  core.NewHandler(),
	}
}

func (c *gactusCore) listenHTTP() error {
	ctx := context.Background()
	logger.Debugf(ctx, "start http server on %s", c.httpAddr)
	return http.ListenAndServe(c.httpAddr, c.handler)
}

func (c *gactusCore) listenTCP() error {
	ctx := context.Background()
	logger.Debugf(ctx, "start tcp server on %s", c.tcpAddr)
	return core.ListenAndServe(c.tcpAddr, c.handler)
}

// Start is used to start core server.
func (c *gactusCore) Start() {
	ctx := context.Background()
	go func() {
		err := c.listenHTTP()
		if err != nil {
			logger.Fatalf(ctx, err.Error())
		}
	}()
	go func() {
		err := c.listenTCP()
		if err != nil {
			logger.Fatalf(ctx, err.Error())
		}
	}()
}

// Wait is used to wait for interrupting signal.
func (c *gactusCore) Wait() {
	p := make(chan os.Signal, 1)
	signal.Notify(p, os.Interrupt, syscall.SIGTERM)
	<-p
	logger.Warnf(context.Background(), "core server is terminated")
	os.Exit(0)
}
