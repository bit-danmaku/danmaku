package handler

import (
	"context"
	"encoding/json"

	log "github.com/asim/go-micro/v3/logger"

	pb "github.com/bit-danmaku/danmaku/proto/kafkaproducer"

	"fmt"

	_ "github.com/asim/go-micro/plugins/broker/kafka/v3"
	"github.com/asim/go-micro/v3/broker"
)

var (
	service = "kafka-producer"
	version = "latest"
)
var (
	topic = "go.micro.topic.foo"
)

type KafkaProducer struct {
}

func InitKafkaProducer() KafkaProducer {
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	return KafkaProducer{}

}

func (kp *KafkaProducer) PostKafka(ctx context.Context, req *pb.PostRequest, rsp *pb.PostResponse) error {

	danmaku := req.Danmaku
	json_danmaku, _ := json.Marshal(danmaku)

	msg := &broker.Message{
		Header: map[string]string{
			"id": fmt.Sprintf("%d", req.ChannelID),
		},
		Body: []byte(fmt.Sprintf("%#v", json_danmaku)),
	}
	if err := broker.Publish(topic, msg); err != nil {
		rsp.Code = 1
		rsp.Msg = err.Error()
		return err
	} else {
		return nil
	}
}
