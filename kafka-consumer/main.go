package main

import (
	"github.com/bit-danmaku/danmaku/kafka-consumer/handler"
	pb "github.com/bit-danmaku/danmaku/proto/kafkaconsumer"
	common "github.com/bit-danmaku/danmaku/common"

	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
)

var (
	version = "latest"
)

func main() {
	hdlStruct := handler.InitKafkaConsumer()
	hdlStruct.Sub()

	//Create service
	srv := micro.NewService(
		micro.Name(common.KAFKA_CONSUMER),
		micro.Version(version),
	)
	srv.Init()

	//Register handler
	pb.RegisterKafkaConsumerHandler(srv.Server(), hdlStruct)
	
	//Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
