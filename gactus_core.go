package gactus

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mr-panta/gactus/pkg/core"
	"github.com/mr-panta/gactus/pkg/tcp"
	"github.com/mr-panta/go-logger"
)

type defaultCore struct {
	httpAddr string
	tcpAddr  string
	handler  *core.Handler
}

// NewCore [TOWRITE]
func NewCore(httpAddr, tcpAddr string) Core {
	return &defaultCore{
		httpAddr: httpAddr,
		tcpAddr:  tcpAddr,
		handler:  core.NewHandler(),
	}
}

func (c *defaultCore) listenHTTP() error {
	ctx := context.Background()
	logger.Debugf(ctx, "start http server at %s", c.httpAddr)
	return http.ListenAndServe(c.httpAddr, c.handler)
}

func (c *defaultCore) listenTCP() error {
	ctx := context.Background()
	logger.Debugf(ctx, "start tcp server at %s", c.tcpAddr)
	return tcp.ListenAndServe(c.tcpAddr, c.handler)
}

// Start is used to start core server.
func (c *defaultCore) Start() {
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
func (c *defaultCore) Wait() {
	p := make(chan os.Signal, 1)
	signal.Notify(p, os.Interrupt, syscall.SIGTERM)
	<-p
	logger.Warnf(context.Background(), "core server is terminated")
	os.Exit(0)
}
