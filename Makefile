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
	go build -o ./bin/core ./cmd/core/*.go
build/example:
	go build -o ./bin/example ./cmd/example/*.go

# run
run/core:
	go run ./cmd/core/*.go
run/example:
	go run ./cmd/example/*.go