package main

import (
	"context"
	"os"

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
		tcpAddr = ":3000"
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

func getProcessorList() []gactus.Processor {
	return []gactus.Processor{&calculateProcessor{}}
}
