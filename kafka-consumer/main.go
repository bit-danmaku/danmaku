package main

import (
	"github.com/bit-danmaku/danmaku/kafka-consumer/handler"
	pb "github.com/bit-danmaku/danmaku/proto/kafkaconsumer"

	"fmt"
	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
)

var (
	service = "kafka-consumer"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterKafkaConsumerHandler(srv.Server(), new(handler.KafkaConsumer))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.print()
}
