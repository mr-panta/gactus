package main

import (
	"context"
	"os"

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
		coreAddr = gactus.DefaultCoreTCPAddr
	}

	// Get TCP address
	tcpAddr := os.Getenv(gactus.ServiceTCPAddrVar)
	if tcpAddr == "" {
		tcpAddr = "127.0.0.1:3000"
	}

	// Setup and start service server
	service, err := gactus.NewService("example", coreAddr, tcpAddr, 1, 10, 100, 10, 1000)
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}
	// Start service
	service.Start()
	// Register processors
	err = service.RegisterProcessors(getProcessorList())
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}

	service.Wait()
}

func getProcessorList() []*gactus.Processor {
	return []*gactus.Processor{
		{
			Command: "example.calculate",
			Req:     &pb.CalculateRequest{},
			Res:     &pb.CalculateResponse{},
			HTTPConfig: &gtpb.HttpConfig{
				Method: gtpb.Constant_HTTP_METHOD_POST,
				Path:   "/calculate",
			},
			Process: func(ctx context.Context, req, res proto.Message) (code uint32) {
				request := req.(*pb.CalculateRequest)
				response := res.(*pb.CalculateResponse)
				response.C = request.A + request.B
				logger.Infof(ctx, "%d + %d = %d", request.A, request.B, response.C)
				return uint32(gtpb.Constant_RESPONSE_OK)
			},
		},
	}
}
