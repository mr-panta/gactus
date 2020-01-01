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
	GO111MODULE=on go build -mod vendor -o ./bin/core ./cmd/core/*.go
build/example:
	GO111MODULE=on go build -mod vendor -o ./bin/example ./cmd/example/*.go

# run
run/core:
	GO111MODULE=on go run -mod vendor ./cmd/core/*.go
run/example:
	GO111MODULE=on go run -mod vendor ./cmd/example/*.go
run/example_2:
	GO111MODULE=on go run -mod vendor ./cmd/example_2/*.go