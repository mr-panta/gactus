package gactus

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/joho/godotenv"
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

func NewService() (Service, error) {
	// Get env
	_ = godotenv.Load()

	// Get service name
	name := os.Getenv(ServiceNameVar)
	if name == "" {
		return nil, errors.New("config GACTUS_SERVICE_NAME does not exist")
	}

	// Get core address
	coreAddr := os.Getenv(ServiceCoreAddrVar)
	if coreAddr == "" {
		return nil, errors.New("config GACTUS_SERVICE_CORE_ADDR does not exist")
	}

	// Get TCP port
	tcpPort, err := strconv.Atoi(os.Getenv(ServiceTCPPortVar))
	if err != nil {
		tcpPort = DefaultServiceTCPPort
	}

	// Get minimum number of connection
	minConns, err := strconv.Atoi(os.Getenv(ServiceMinConnsVar))
	if err != nil {
		minConns = DefaultServiceMinConns
	}

	// Get maximum number of connection
	maxConns, err := strconv.Atoi(os.Getenv(ServiceMaxConnsVar))
	if err != nil {
		maxConns = DefaultServiceMaxConns
	}

	// Get duration idle connection timeout
	idleConnTimeout, err := strconv.Atoi(os.Getenv(ServiceIdleConnTimeoutVar))
	if err != nil {
		idleConnTimeout = DefaultServiceIdleConnTimeout
	}

	// Get duration waiting connection timeout
	waitConnTimeout, err := strconv.Atoi(os.Getenv(ServiceWaitConnTimeoutVar))
	if err != nil {
		waitConnTimeout = DefaultServiceWaitConnTimeout
	}

	// Get duration clearing connection pool
	clearPeriod, err := strconv.Atoi(os.Getenv(ServiceClearPeriodVar))
	if err != nil {
		clearPeriod = DefaultServiceClearPeriod
	}

	return NewServiceWithConfig(
		name,
		coreAddr,
		tcpPort,
		minConns,
		maxConns,
		idleConnTimeout,
		waitConnTimeout,
		clearPeriod,
	)
}

func NewServiceWithConfig(name, coreAddr string, tcpPort, minConns, maxConns, idleConnTimeout, waitConnTimeout,
	clearPeriod int) (Service, error) {

	ctx := context.Background()
	logger.Infof(ctx, "GACTUS_SERVICE_NAME=%s", name)
	logger.Infof(ctx, "GACTUS_SERVICE_CORE_ADDR=%s", coreAddr)
	logger.Infof(ctx, "GACTUS_SERVICE_TCP_PORT=%d", tcpPort)
	logger.Infof(ctx, "GACTUS_SERVICE_MIN_CONNS=%d", minConns)
	logger.Infof(ctx, "GACTUS_SERVICE_MAX_CONNS=%d", maxConns)
	logger.Infof(ctx, "GACTUS_SERVICE_IDLE_CONN_TIMEOUT=%d", idleConnTimeout)
	logger.Infof(ctx, "GACTUS_SERVICE_WAIT_CONN_TIMEOUT=%d", waitConnTimeout)
	logger.Infof(ctx, "GACTUS_SERVICE_CLEAR_PERIOD=%d", clearPeriod)

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
