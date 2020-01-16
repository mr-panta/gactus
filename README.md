# Gactus - Go Microservice Framework

Gactus is a microservice framework library for Go. It provides ability to replicate instances for your microservice and ability to communicate between microservices via basic RPC communication but still be able to publish APIs via HTTP.

## Installation

To apply Gactus to your application, there are two parts you need to install in your system. First, **Gactus Core**  is the part that provides HTTP communication for clients, contains service registries, and redirects HTTP request to other services. Second, **Gactus Service** is the part that contains all business logic, communicates with other Gactus Services and Gactus Core.
