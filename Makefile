# gen
gen:
	make gen/proto
gen/proto:
	cd proto; protoc --go_out=../ gactus.proto; cd ..

# build
build:
	make build/core
build/core:
	go build -o ./bin/core ./cmd/core/*.go
