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
        "gactus service started with name=%s on tcp port=%s and connect to gactus core address=%s with access key=%s",
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

### Call Gactus API

### Expose Gactus API to HTTP

### Upload file with Gactus API
