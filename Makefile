GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v3@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=./proto/kafkaproducer/ --go_out=./proto/kafkaproducer/ ./proto/kafkaproducer/kafka-producer.proto
	@protoc --proto_path=. --micro_out=./proto/kafkaconsumer/ --go_out=:./proto/kafkaconsumer/ ./proto/kafkaconsumer/kafka-consumer.proto
	@protoc --proto_path=. --micro_out=./proto/danmakucache/ --go_out=:./proto/danmakucache/ -I=./proto ./proto/danmakucache/danmaku-cache.proto
	@protoc --proto_path=. --micro_out=paths=source_relative:. --go_out=paths=source_relative:. ./proto/common/danmaku.proto


.PHONY: services
services: api-gateway kafka-producer kafka-consumer danmaku-cache

.PHONY: api-gateway
api-gateway:
	@go build -o bin/api-gateway ./api-gateway-client/main.go

.PHONY: kafka-producer
kafka-producer:
	@go build -o bin/kafka-producer ./kafka-producer/main.go

.PHONY: kafka-consumer
kafka-consumer:
	@go build -o bin/kafka-consumer ./kafka-consumer/main.go

.PHONY: danmaku-cache
danmaku-cache:
	@go build -o bin/danmaku-cache ./danmaku-cache/main.go

.PHONY: tidy
tidy:
	@go mod tidy

#.PHONY: test
#test:
	#@go test -v ./... -cover

#.PHONY: docker
#docker:
	#@docker build -t kafka-consumer:latest .
