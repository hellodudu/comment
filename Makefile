
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build main.go

.PHONY: proto
proto:
	protoc -I=proto --go_out=plugins=grpc:proto proto/world_message.proto

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t ultimate .
