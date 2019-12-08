package gactus

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/mr-panta/gactus/internal/core"
	"github.com/mr-panta/go-logger"
)

type gactusCore struct {
	httpPort int
	tcpPort  int
	handler  core.Handler
}

// NewCore [TOWRITE]
func NewCore() Core {
	// Get env
	_ = godotenv.Load()

	// Get HTTP port
	httpPort, err := strconv.Atoi(os.Getenv(CoreHTTPPortVar))
	if err != nil {
		httpPort = DefaultCoreHTTPPort
	}

	// Get TCP port
	tcpPort, err := strconv.Atoi(os.Getenv(CoreTCPPortVar))
	if err != nil {
		tcpPort = DefaultCoreTCPPort
	}

	// Get health check interval
	healthCheckInterval, err := strconv.Atoi(os.Getenv(CoreDefaultHealthCheckIntervalVar))
	if err != nil {
		healthCheckInterval = DefaultHealthCheckInterval
	}

	// Get access token
	accessToken := os.Getenv(CoreAccessToken)

	return NewCoreWithConfig(httpPort, tcpPort, accessToken, healthCheckInterval)
}

func NewCoreWithConfig(httpPort, tcpPort int, accessToken string, healthCheckInterval int) Core {
	ctx := context.Background()
	logger.Infof(ctx, "GACTUS_CORE_HTTP_PORT=%d", httpPort)
	logger.Infof(ctx, "GACTUS_CORE_TCP_PORT=%d", tcpPort)
	logger.Infof(ctx, "GACTUS_CORE_HEALTH_CHECK_INTERVAL=%d", healthCheckInterval)
	logger.Infof(ctx, "GACTUS_CORE_ACCESS_TOKEN=%s", accessToken)

	return &gactusCore{
		httpPort: httpPort,
		tcpPort:  tcpPort,
		handler:  core.NewHandler(accessToken, healthCheckInterval),
	}
}

func (c *gactusCore) listenHTTP() error {
	ctx := context.Background()
	logger.Debugf(ctx, "start http server on port %d", c.httpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", c.httpPort), c.handler)
}

func (c *gactusCore) listenTCP() error {
	ctx := context.Background()
	logger.Debugf(ctx, "start tcp server on port %d", c.tcpPort)
	return core.ListenAndServe(fmt.Sprintf(":%d", c.tcpPort), c.handler)
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
