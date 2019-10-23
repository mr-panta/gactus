package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/cmd/example/proto"
	gtpb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"

	"github.com/mr-panta/gactus"
)

func main() {
	ctx := context.Background()

	// Get core address
	coreAddr := os.Getenv(gactus.ServiceCoreAddrVar)
	if coreAddr == "" {
		coreAddr = fmt.Sprintf("127.0.0.1:%d", gactus.DefaultCoreTCPPort)
	}

	// Get TCP port
	tcpPort, err := strconv.Atoi(os.Getenv(gactus.ServiceTCPPortVar))
	if err != nil {
		tcpPort = 3000
	}

	// Setup and start service server
	service, err := gactus.NewService("calculator", coreAddr, tcpPort, 1, 10, 100, 10, 1000)
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}
	// Start service
	service.Start()
	// Register processors
	err = service.RegisterService(getProcessorList())
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}

	service.Wait()
}

func getProcessorList() []*gactus.Processor {
	return []*gactus.Processor{
		{
			Command: "calculator.add",
			Req:     &pb.CalculateRequest{},
			Res:     &pb.CalculateResponse{},
			HTTPConfig: &gtpb.HttpConfig{
				Method: gtpb.Constant_HTTP_METHOD_POST,
				Path:   "/calculator/add",
			},
			HTTPMiddleware: func(ctx context.Context, header map[string]string, req, res proto.Message) {
				for key, value := range header {
					logger.Debugf(ctx, "%s: %s", key, value)
				}
			},
			Process: func(ctx context.Context, req, res proto.Message) (code uint32) {
				request := req.(*pb.CalculateRequest)
				response := res.(*pb.CalculateResponse)
				response.C = request.A + request.B
				logger.Debugf(ctx, "%d", len(request.Files))
				for _, f := range request.Files {
					logger.Debugf(ctx, "name:%s | size:%d bytes", string(f.Name), len(f.Content))
				}
				return uint32(gtpb.Constant_RESPONSE_OK)
			},
		},
		{
			Command: "calculator.substract",
			Req:     &pb.CalculateRequest{},
			Res:     &pb.CalculateResponse{},
			HTTPConfig: &gtpb.HttpConfig{
				Method: gtpb.Constant_HTTP_METHOD_POST,
				Path:   "/calculator/substract",
			},
			Process: func(ctx context.Context, req, res proto.Message) (code uint32) {
				request := req.(*pb.CalculateRequest)
				response := res.(*pb.CalculateResponse)
				response.C = request.A - request.B
				return uint32(gtpb.Constant_RESPONSE_OK)
			},
		},
	}
}
