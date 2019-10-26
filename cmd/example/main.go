package main

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/joho/godotenv"
	pb "github.com/mr-panta/gactus/cmd/example/proto"
	gtpb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"

	"github.com/mr-panta/gactus"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load("./cmd/example/.env")
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}
	// Setup and start service server
	service, err := gactus.NewService()
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
			HTTPMiddleware: func(ctx context.Context, header, query map[string]string, req, res proto.Message) (code uint32) {
				for key, value := range header {
					logger.Debugf(ctx, "%s: %s", key, value)
				}
				return 0
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
