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
build/core:
	go build -mod vendor -o ./bin/core ./cmd/core/*.go
build/example:
	go build -mod vendor -o ./bin/example ./cmd/example/*.go

# run
run/core:
	go run -mod vendor ./cmd/core/*.go
run/example:
	go run -mod vendor ./cmd/example/*.go