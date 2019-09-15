package core

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mr-panta/gactus/pkg/logger"
	"github.com/mr-panta/gactus/pkg/tcp"
)

// Core contains core server config and provides methods for creating core server.
type Core struct {
	httpAddr string
	tcpAddr  string
	handler  *handler
}

type commandData struct {
	command string
	service string
}

// NewCore is used to create NewCore.
func NewCore(httpAddr, tcpAddr string) *Core {
	return &Core{
		httpAddr: httpAddr,
		tcpAddr:  tcpAddr,
	}
}

// Start is used to start core server.
func (c *Core) Start() {
	ctx := context.Background()
	c.handler = newHandler()
	// c.handler.serviceManager.routeToCommandMap["POST:/login"] = "user.login"
	go func() {
		logger.Infof(ctx, "start HTTP server at %s", c.httpAddr)
		err := c.listenHTTP()
		if err != nil {
			logger.Fatalf(ctx, err.Error())
		}
	}()
	go func() {
		logger.Infof(ctx, "start TCP server at %s", c.tcpAddr)
		err := c.listenTCP()
		if err != nil {
			logger.Fatalf(ctx, err.Error())
		}
	}()
}

// Wait is used to wait for interrupting signal.
func (c *Core) Wait() {
	p := make(chan os.Signal, 1)
	signal.Notify(p, os.Interrupt, syscall.SIGTERM)
	<-p
	logger.Warnf(context.Background(), "core server is terminated")
	os.Exit(0)
}

func (c *Core) listenHTTP() error {
	return http.ListenAndServe(c.httpAddr, c.handler)
}

func (c *Core) listenTCP() error {
	return tcp.ListenAndServe(c.tcpAddr, c.handler)
}
