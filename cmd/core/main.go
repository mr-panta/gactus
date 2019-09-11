package main

import (
	"os"

	"github.com/mr-panta/gactus/pkg/core"
)

func main() {
	// Get HTTP address
	httpAddr := os.Getenv(core.HTTPAddrVar)
	if httpAddr == "" {
		httpAddr = core.DefaultHTTPAddr
	}

	// Get TCP address
	tcpAddr := os.Getenv(core.TCPAddrVar)
	if tcpAddr == "" {
		tcpAddr = core.DefaultTCPAddr
	}

	// Setup and start core server
	server := core.NewCore(httpAddr, tcpAddr)
	server.Start()
	server.Wait()
}
