package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mr-panta/gactus/service"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus"
	"github.com/mr-panta/gactus/cmd/example/echo"
)

func main() {
	s1 := gactus.NewGactusService(3000)
	s1.RegisterProcessors([]*gactus.Processor{
		{
			Command: "echo.do",
			Req:     &echo.Request{},
			Res:     &echo.Response{},
			Process: func(ctx context.Context, request, response proto.Message) error {
				req := request.(*echo.Request)
				res := response.(*echo.Response)
				res.Message = req.Message + ", Okay"
				return nil
			},
		},
	})
	go func() {
		err := s1.Start()
		if err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(time.Second)

	s2 := service.NewService()
	err := s2.SetAddressCommands(":3000", []string{"echo.do"})
	if err != nil {
		fmt.Println(err)
	}
	req := &echo.Request{Message: "Hi"}
	res := &echo.Response{}
	err = s2.SendRequest(context.Background(), "echo.do", req, res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)
}
