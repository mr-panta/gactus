package main

import (
	"os"

	"github.com/mr-panta/gactus"
)

func main() {
	// Get HTTP address
	httpAddr := os.Getenv(gactus.CoreHTTPAddrVar)
	if httpAddr == "" {
		httpAddr = gactus.DefaultCoreHTTPAddr
	}

	// Get TCP address
	tcpAddr := os.Getenv(gactus.CoreTCPAddrVar)
	if tcpAddr == "" {
		tcpAddr = gactus.DefaultCoreTCPAddr
	}

	// Setup and start core server
	core := gactus.NewCore(httpAddr, tcpAddr)
	core.Start()
	core.Wait()
}
