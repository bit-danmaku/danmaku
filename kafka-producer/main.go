package main

import (
	"github.com/bit-danmaku/danmaku/kafka-producer/handler"
	pb "github.com/bit-danmaku/danmaku/proto/kafkaproducer"
	common "github.com/bit-danmaku/danmaku/common"

	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
)

var (
	version = "latest"
)

func main() {
	hdlStruct := handler.InitKafkaProducer()
	// Create service
	srv := micro.NewService(
		micro.Name(common.KAFKA_PRODUCER),
		micro.Version(version),
	)
	srv.Init()
	// Register handler
	pb.RegisterKafkaProducerHandler(srv.Server(), hdlStruct)

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
