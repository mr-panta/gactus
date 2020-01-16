# Gactus - Go Microservice Framework

Gactus is a microservice framework library for Go. It provides ability to replicate instances for your microservice and ability to communicate between microservices via basic RPC communication but still be able to publish APIs via HTTP.

- [Installation](#installation)
  - [Gactus Core](#gactus-core)
  - [Gactus Service](#gactus-service)
- [Usage](#usage)
  - [Create Gactus API](#create-gactus-api)
  - [Call Gactus API](#call-gactus-api)
  - [Expose Gactus API to HTTP](#expose-gactus-api-to-http)
  - [Upload file with Gactus API](#upload-file-with-api)

## Installation

To apply Gactus to your application, there are two parts you need to install in your system. First, **Gactus Core** is the part that provides HTTP communication for clients, contains service registries, and redirects HTTP request to other services. Second, **Gactus Service** is the part that contains all business logic, communicates with other Gactus Services and Gactus Core.

![General system diagram with Gactus](https://raw.githubusercontent.com/mr-panta/gactus/feature/init/doc/gactus.png?raw=true)

### Gactus Core

Core provides HTTP communication with clients, contains service registries and communicates with other Gactus services via RPC. Below is how this part is installed.

```go
package main

import (
    "github.com/mr-panta/gactus"
    "github.com/mr-panta/go-logger"
)

func main() {
    httpPort := 80            // To receive HTTP request
    tcpPort := 3000           // To receive RPC request
    accessKey := "secret1234" // For authentication
    core := gactus.NewGactusCore(httpPort, tcpPort, accessKey)
    core.Start()
    logger.Infof(
        context.Background(),
        "gactus core started on http port=%d, tcp port=%d with access key=%s ",
        httpPort,
        tcpPort,
        accessKey,
    )
    core.Wait()
}
```

With only access key for validating Gactus Services that are going to connect with the core, HTTP port for allowing clients to connect, and TCP port for RPC communication, the Gactus Core can be installed easily.

### Gactus Service

Gactus Service is the part that processes all of your business logic and can send requests to other services as well. To install this part you can follow the code below.

```go
package main

import (
    "github.com/mr-panta/gactus"
    "github.com/mr-panta/go-logger"
)

func main() {
    serviceName := "example"
    tcpPort := 4000                   // To receive RPC request
    coreAddress := "196.168.0.2:3000" // address of gactus core
    accessKey := "secret1234"         // same as the one in Gactus Core
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
    service.RegisterProcessors([]*gactus.Processor{}) // TODO: register processors for providing APIs
    service.Wait()
}
```

The code above is the code to show you an example of how to install Gactus Service in your system, but this is the Gactus Service installation, you still need to create APIs and registers them to Gactus Core. And you will find these steps below.

## Usage

### Create Gactus API

Before creating your API, you need to define the protocol you will use with the API. In Gactus, you need to create request and response with [Protocol Buffers](https://developers.google.com/protocol-buffers). We also recomment you to use [proto3](https://developers.google.com/protocol-buffers/docs/proto3) with Gactus.

```proto3
syntax = "proto3";
package first_example;

message AddRequest {
    uint32 a = 1;
    uint32 b = 2;
}

message AddResponse {
    uint32 c = 1;
}
```

Gactus library provides `gactus.Processor` which is a struct used to describe your API. Below is the example code of how to create your own processor.

```go
service := gactus.NewGactusService(
    "first_example",
    tcpPort,
    coreAddress,
    accessKey,
)
service.Start()
processors := []*gactus.Processor{
    {
        Command: "first_example.add",
        Req:     &first_example.AddRequest{},
        Res:     &first_example.AddResponse{},
        Process: func(ctx context.Context, request, response proto.Message) error {
            req, ok := request.(*first_example.AddRequest)
            if !ok {
                return errors.New("cannot assert request object")
            }
            res, ok := response.(*first_example.AddResponse)
            if !ok {
                return errors.New("cannot assert response object")
            }
            res.C = req.A + req.B
            return nil
        },
    },
}
err := service.RegisterProcessors(processors)
if err != nil {
    logger.Fatalf(context.Background(), err.Error())
}
service.Wait()
```

From the example, you will get the service with an API for doing some basic calculation.

### Call Gactus API

After you can have your API, sometimes the API need to be called by other API. There is a method in Gactus Service object called `SendRequest` that you can use to call other services APIs.

```proto3
syntax = "proto3";
package second_example;

message SubtractRequest {
    uint32 a = 1;
    uint32 b = 2;
}

message SubtractResponse {
    uint32 c = 1;
}
```

```go
service := gactus.NewGactusService(
    "second_example",
    tcpPort,
    coreAddress,
    accessKey,
)
service.Start()
processors := []*gactus.Processor{
    {
        Command: "second_example.subtract",
        Req:     &second_example.SubtractRequest{},
        Res:     &second_example.SubtractResponse{},
        Process: func(ctx context.Context, request, response proto.Message) error {
            req, ok := request.(*second_example.SubtractRequest)
            if !ok {
                return errors.New("cannot assert request object")
            }
            res, ok := response.(*second_example.SubtractResponse)
            if !ok {
                return errors.New("cannot assert response object")
            }
            addReq := &first_example.AddRequest{
                A: req.A,
                B: -req.B,
            }
            addRes := &first_example.AddResponse{}
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
    logger.Fatalf(context.Background(), err.Error())
}
service.Wait()
```

### Expose Gactus API to HTTP

// TODO

### Upload file with Gactus API

// TODO
