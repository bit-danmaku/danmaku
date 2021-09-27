package main

import (
	"github.com/bit-danmaku/danmaku/kafka-producer/handler"
	pb "github.com/bit-danmaku/danmaku/proto/kafkaproducer"

	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
)

var (
	service = "kafka-producer"
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
	pb.RegisterKafkaProducerHandler(srv.Server(), new(handler.KafkaProducer))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
