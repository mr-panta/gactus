package main

import (
	"github.com/mr-panta/gactus/pkg/core"
)

func main() {
	server := core.NewCore(":8000", ":8001")
	server.Start()
	server.Wait()
}
