package main

import (
	"os"
	"strconv"

	"github.com/mr-panta/gactus"
)

func main() {
	// Get HTTP port
	httpPort, err := strconv.Atoi(os.Getenv(gactus.CoreHTTPPortVar))
	if err != nil {
		httpPort = gactus.DefaultCoreHTTPPort
	}

	// Get TCP port
	tcpPort, err := strconv.Atoi(os.Getenv(gactus.CoreTCPPortVar))
	if err != nil {
		tcpPort = gactus.DefaultCoreTCPPort
	}

	// Get health check interval
	healthCheckInterval, err := strconv.Atoi(os.Getenv(gactus.CoreDefaultHealthCheckIntervalVar))
	if err != nil {
		healthCheckInterval = gactus.DefaultHealthCheckInterval
	}

	core := gactus.NewCore(httpPort, tcpPort, healthCheckInterval)
	core.Start()
	core.Wait()
}
