package handler

import (
	"context"
	"encoding/json"

	log "github.com/asim/go-micro/v3/logger"

	pb "github.com/bit-danmaku/danmaku/proto/kafkaproducer"
	common "github.com/bit-danmaku/danmaku/common"

	"fmt"

	kafka "github.com/asim/go-micro/plugins/broker/kafka/v3"
	"github.com/asim/go-micro/v3/broker"
)

var (
	version = "latest"
)

type KafkaProducer struct {
	kafkaBroker broker.Broker
}

func InitKafkaProducer() *KafkaProducer {
	kafkaBroker := kafka.NewBroker()
	if err := kafkaBroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := kafkaBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	return &KafkaProducer{kafkaBroker: kafkaBroker}
}

func (kp *KafkaProducer) PostKafka(ctx context.Context, req *pb.PostRequest, rsp *pb.PostResponse) error {
	//log.Infof("Received KafkaProducer.PostRequest request: %+v", req)
	danmaku := req.Danmaku
	json_danmaku, _ := json.Marshal(danmaku)

	msg := &broker.Message{
		Header: map[string]string{
			"id": fmt.Sprintf("%d", req.ChannelID),
		},
		Body: []byte(fmt.Sprintf("%s", json_danmaku)),
	}
	if err := kp.kafkaBroker.Publish(common.TOPIC, msg); err != nil {
		rsp.Code = 1
		rsp.Msg = err.Error()
		return err
	} else {
		return nil
	}
}
