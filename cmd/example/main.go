package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus"
	"github.com/mr-panta/gactus/cmd/example/example"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
)

func main() {
	runCore()
	runFirst()
	runSecond()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func runCore() {
	httpPort := 80            // To receive HTTP request
	tcpPort := 3000           // To receive RPC request
	accessKey := "secret1234" // For authorization
	core := gactus.NewGactusCore(httpPort, tcpPort, accessKey)
	core.Start()
	logger.Infof(
		context.Background(),
		"gactus core started on http port=%d, tcp port=%d with access key=%s ",
		httpPort,
		tcpPort,
		accessKey,
	)
}

func runFirst() {
	serviceName := "first_example"
	tcpPort := 4000                 // To receive RPC request
	coreAddress := "localhost:3000" // address of gactus core
	accessKey := "secret1234"       // same as the one in Gactus Core
	service := gactus.NewGactusService(
		serviceName,
		tcpPort,
		coreAddress,
		accessKey,
	)
	service.Start()
	logger.Infof(
		context.Background(),
		"gactus service started with name=%s on tcp port=%d and connect to gactus core address=%s with access key=%s",
		serviceName,
		tcpPort,
		coreAddress,
		accessKey,
	)
	processors := []*gactus.Processor{
		{
			Command: "first_example.add",
			Req:     &example.AddRequest{},
			Res:     &example.AddResponse{},
			HTTPConfig: &pb.HttpConfig{
				Method: pb.Constant_HTTP_METHOD_GET,
				Path:   "/first-example/add",
			},
			HTTPMiddleware: func(ctx context.Context, header, query map[string]string, request, response proto.Message) error {
				req, ok := request.(*example.AddRequest)
				if !ok {
					return errors.New("cannot assert request object")
				}
				a, _ := strconv.ParseInt(query["a"], 10, 32)
				b, _ := strconv.ParseInt(query["b"], 10, 32)
				req.A = int32(a)
				req.B = int32(b)
				return nil
			},
			Process: func(ctx context.Context, request, response proto.Message) error {
				req, ok := request.(*example.AddRequest)
				if !ok {
					return errors.New("cannot assert request object")
				}
				res, ok := response.(*example.AddResponse)
				if !ok {
					return errors.New("cannot assert response object")
				}
				res.C = req.A + req.B
				return nil
			},
		},
		{
			Command: "first_example.change_profile",
			Req:     &example.ChangeProfileRequest{},
			Res:     &example.ChangeProfileResponse{},
			HTTPConfig: &pb.HttpConfig{
				Method: pb.Constant_HTTP_METHOD_POST,
				Path:   "/first-example/change-profile",
			},
			Process: func(ctx context.Context, request, response proto.Message) error {
				req, ok := request.(*example.ChangeProfileRequest)
				if !ok {
					return errors.New("cannot assert request object")
				}
				res, ok := response.(*example.ChangeProfileResponse)
				if !ok {
					return errors.New("cannot assert response object")
				}
				logger.Infof(ctx, "name=%s", req.Picture.Name)
				res.FileSize = uint32(len(req.Picture.Content))
				return nil
			},
		},
	}
	err := service.RegisterProcessors(processors)
	if err != nil {
		logger.Errorf(context.Background(), err.Error())
	}
}

func runSecond() {
	serviceName := "second_example"
	tcpPort := 4001                 // To receive RPC request
	coreAddress := "localhost:3000" // address of gactus core
	accessKey := "secret1234"       // same as the one in Gactus Core
	service := gactus.NewGactusService(
		serviceName,
		tcpPort,
		coreAddress,
		accessKey,
	)
	service.Start()
	logger.Infof(
		context.Background(),
		"gactus service started with name=%s on tcp port=%d and connect to gactus core address=%s with access key=%s",
		serviceName,
		tcpPort,
		coreAddress,
		accessKey,
	)
	processors := []*gactus.Processor{
		{
			Command: "second_example.subtract",
			Req:     &example.SubtractRequest{},
			Res:     &example.SubtractResponse{},
			HTTPConfig: &pb.HttpConfig{
				Method: pb.Constant_HTTP_METHOD_POST,
				Path:   "/second-example/subtract",
			},
			Process: func(ctx context.Context, request, response proto.Message) error {
				req, ok := request.(*example.SubtractRequest)
				if !ok {
					return errors.New("cannot assert request object")
				}
				res, ok := response.(*example.SubtractResponse)
				if !ok {
					return errors.New("cannot assert response object")
				}
				addReq := &example.AddRequest{
					A: req.A,
					B: -req.B,
				}
				addRes := &example.AddResponse{}
				err := service.SendRequest(ctx, "first_example.add", addReq, addRes)
				if err != nil {
					return fmt.Errorf("fail to call first_example.add, err=%v", err)
				}
				res.C = addRes.C
				return nil
			},
		},
	}
	err := service.RegisterProcessors(processors)
	if err != nil {
		logger.Errorf(context.Background(), err.Error())
	}
}
