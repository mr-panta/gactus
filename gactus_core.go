package gactus

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	bd "github.com/mr-panta/gactus/body"
	"github.com/mr-panta/gactus/config"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/gactus/service"
	"github.com/mr-panta/go-logger"
)

type GactusCore interface {
	Start()
	Wait()
}

type gactusCore struct {
	lock            sync.RWMutex
	tcpPort         int
	httpPort        int
	service         *service.Service
	routeCommandMap map[string]string
}

func NewGactusCore(httpPort, tcpPort int) GactusCore {
	gc := &gactusCore{
		tcpPort:         tcpPort,
		httpPort:        httpPort,
		service:         service.NewService(),
		routeCommandMap: make(map[string]string),
	}
	gc.service.AddProcessor(config.CMDCoreRegisterService, &service.Processor{
		Req:     &pb.RegisterServiceRequest{},
		Res:     &pb.RegisterServiceResponse{},
		Process: gc.ProcessRegisterService,
	})
	return gc
}

func getRoute(method, path string) string {
	return fmt.Sprintf("%s:%s", strings.ToUpper(method), path)
}

func getMethodString(method pb.Constant_HttpMethod) string {
	switch method {
	case pb.Constant_HTTP_METHOD_GET:
		return "GET"
	case pb.Constant_HTTP_METHOD_POST:
		return "POST"
	}
	return "UNKNOWN"
}

func (gc *gactusCore) Start() {
	ctx := logger.GetContextWithLogID(context.Background(), "start_core")
	go gc.startService(ctx)
	go gc.startHTTPServer(ctx)
}

func (gc *gactusCore) Wait() {
	p := make(chan os.Signal, 1)
	signal.Notify(p, os.Interrupt, syscall.SIGTERM)
	<-p
	logger.Warnf(context.Background(), "core server is terminated")
	os.Exit(0)
}

func (gc *gactusCore) startService(ctx context.Context) {
	lis, err := net.Listen(networkTCP, fmt.Sprintf(":%d", gc.tcpPort))
	if err != nil {
		logger.Fatalf(ctx, "cannot listen tcp to port=%d, err=%v", gc.tcpPort, err)
	}
	server := rpc.NewServer()
	err = server.Register(gc.service)
	if err != nil {
		logger.Fatalf(ctx, "cannot register rpc receiver, err=%v", err)
	}
	logger.Infof(ctx, "start tcp server, port=%d", gc.tcpPort)
	server.Accept(lis)
}

func (gc *gactusCore) startHTTPServer(ctx context.Context) {
	logger.Infof(ctx, "start http server, port=%d", gc.httpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", gc.httpPort), gc)
	if err != nil {
		logger.Fatalf(ctx, "cannot listen and server http servier, err=%v", err)
	}
}

func generateLogID(ctx context.Context, method, path string) (coveredCTX context.Context, logID string) {
	if ctx == nil {
		ctx = context.Background()
	}
	coveredCTX = logger.GetContextWithLogID(ctx, fmt.Sprintf("%s_%s", method, path))
	return coveredCTX, logger.GetLogID(coveredCTX)
}

func (gc *gactusCore) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	ctx, logID := generateLogID(ctx, req.Method, req.URL.Path)

	body := []byte{}
	statusCode := http.StatusOK
	wrappedRes := &pb.Response{}
	var err error

	defer func() {
		// Send response back
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(statusCode)
		if err != nil {
			logger.Errorf(ctx, err.Error())
			body, _ = bd.Marshal(&pb.ErrorResponse{
				DebugMessage: err.Error(),
			})
		}
		_, _ = res.Write(body)
	}()

	// Get content type
	contentType, rawContentType, err := bd.GetContentTypeValue(req.Header)
	if err != nil {
		err = fmt.Errorf("cannot convert content-type, err=%v", err)
		statusCode = http.StatusInternalServerError
		return
	}

	// Get command from method and path
	command, exists := gc.getCommand(req.Method, req.URL.Path)
	if !exists {
		err = fmt.Errorf("%s:%s route not found", req.Method, req.URL.Path)
		statusCode = http.StatusNotFound
		return
	}

	// Get header
	header := make(map[string]string)
	for key, values := range req.Header {
		if len(values) > 0 {
			header[key] = values[0]
		}
	}

	// Get query
	query := make(map[string]string)
	for key, values := range req.URL.Query() {
		if len(values) > 0 {
			query[key] = values[0]
		}
	}

	// Get body
	if contentType != pb.Constant_CONTENT_TYPE_UNKNOWN {
		body, err = ioutil.ReadAll(req.Body)
		if err != nil {
			err = fmt.Errorf("cannot read body, err=%v", err)
			statusCode = http.StatusInternalServerError
			return
		}
	}

	// Get sender address
	httpAddr := strings.Split(req.RemoteAddr, ":")[0]

	// Setup gactus request
	wrappedReq := &pb.Request{
		HttpAddress:    httpAddr,
		Command:        command,
		LogId:          logID,
		ContentType:    contentType,
		RawContentType: rawContentType,
		Header:         header,
		Query:          query,
		Body:           body,
		IsProto:        false,
	}

	// Send data to TCP
	err = gc.service.SendWrappedRequest("", wrappedReq, wrappedRes)
	if err != nil {
		err = fmt.Errorf("cannot send data to service, err=%v", err)
		statusCode = http.StatusInternalServerError
		return
	}

	// Unwrap response
	body = wrappedRes.Body
}

func (gc *gactusCore) getCommand(method, path string) (command string, exists bool) {
	route := getRoute(method, path)
	return gc.getCommandByRoute(route)
}

func (gc *gactusCore) getCommandByRoute(route string) (command string, exists bool) {
	gc.lock.RLock()
	defer gc.lock.RUnlock()
	command, exists = gc.routeCommandMap[route]
	return command, exists
}

func (gc *gactusCore) setCommandByRoute(route, command string) {
	gc.lock.Lock()
	defer gc.lock.Unlock()
	gc.routeCommandMap[route] = command
}
