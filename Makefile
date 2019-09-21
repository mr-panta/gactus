# install
install:
	go get -u
	make setup

# setup
setup:
	GO111MODULE=on go mod vendor

# gen
gen:
	make gen/proto
gen/proto:
	protoc --go_out=./ ./proto/gactus.proto

# build
build:
	make build/core
	make build/example
	make build/example-2
build/core:
	go build -mod vendor -o ./bin/core ./cmd/core/*.go
build/example:
	go build -mod vendor -o ./bin/example ./cmd/example/*.go
build/example-2:
	go build -mod vendor -o ./bin/example-2 ./cmd/example-2/*.go

# run
run/core:
	go run -mod vendor ./cmd/core/*.go
run/example:
	go run -mod vendor ./cmd/example/*.go
run/example-2:
	go run -mod vendor ./cmd/example-2/*.go