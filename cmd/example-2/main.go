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

	// Get TCP address
	tcpAddr, err := strconv.Atoi(os.Getenv(gactus.ServiceTCPPortVar))
	if err != nil {
		tcpAddr = 4000
	}

	// Setup and start service server
	service, err := gactus.NewService("calculator-2", coreAddr, tcpAddr, 1, 10, 100, 10, 1000)
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}
	// Start service
	service.Start()
	// Register processors
	err = service.RegisterService(getProcessorList(service))
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}

	service.Wait()
}

func getProcessorList(service gactus.Service) []*gactus.Processor {
	return []*gactus.Processor{
		{
			Command: "calculator.multiple",
			Req:     &pb.CalculateRequest{},
			Res:     &pb.CalculateResponse{},
			HTTPConfig: &gtpb.HttpConfig{
				Method: gtpb.Constant_HTTP_METHOD_POST,
				Path:   "/calculator/multiple",
			},
			Process: func(ctx context.Context, req, res proto.Message) (code uint32) {
				request := req.(*pb.CalculateRequest)
				response := res.(*pb.CalculateResponse)
				response.C = 0
				for i := 0; i < int(request.B); i++ {
					calReq := &pb.CalculateRequest{
						A: response.C,
						B: request.A,
					}
					calRes := &pb.CalculateResponse{}
					code := service.SendRequest(context.Background(), "calculator.add", calReq, calRes)
					logger.Debugf(context.Background(), "%d", code)
					logger.Debugf(context.Background(), "%v", calReq)
					logger.Debugf(context.Background(), "%v", calRes)
					response.C = calRes.C
				}
				return uint32(gtpb.Constant_RESPONSE_OK)
			},
		},
		{
			Command: "calculator.divide",
			Req:     &pb.CalculateRequest{},
			Res:     &pb.CalculateResponse{},
			HTTPConfig: &gtpb.HttpConfig{
				Method: gtpb.Constant_HTTP_METHOD_POST,
				Path:   "/calculator/divide",
			},
			Process: func(ctx context.Context, req, res proto.Message) (code uint32) {
				request := req.(*pb.CalculateRequest)
				response := res.(*pb.CalculateResponse)
				response.C = request.A / request.B
				return uint32(gtpb.Constant_RESPONSE_OK)
			},
		},
	}
}
