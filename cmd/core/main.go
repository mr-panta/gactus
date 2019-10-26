package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/mr-panta/gactus"
	"github.com/mr-panta/go-logger"
)

func main() {
	err := godotenv.Load("./cmd/core/.env")
	if err != nil {
		logger.Fatalf(context.Background(), err.Error())
	}
	core := gactus.NewCore()
	core.Start()
	core.Wait()
}
