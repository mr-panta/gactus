package main

import (
	"context"
	"os"

	"github.com/mr-panta/gactus/pkg/logger"

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
		tcpAddr = ":3000"
	}

	// Setup and start service server
	service := gactus.NewService(coreAddr, tcpAddr, 100)
	err := service.RegisterProcessors(nil)
	if err != nil {
		logger.Fatalf(ctx, err.Error())
	}
	service.Start()
	service.Wait()
}
