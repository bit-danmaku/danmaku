package handler

import (
	"context"
	"encoding/json"

	log "github.com/asim/go-micro/v3/logger"

	common "github.com/bit-danmaku/danmaku/common"
	pb "github.com/bit-danmaku/danmaku/proto/kafkaproducer"

	"fmt"

	_ "github.com/asim/go-micro/plugins/broker/kafka/v3"
	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/v3/cmd"
)

var (
	version = "latest"
)

type KafkaProducer struct {
	//kafkaBroker broker.Broker
}

func InitKafkaProducer() *KafkaProducer {
	cmd.Init()
	//kafkaBroker := broker.NewBroker()
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	return &KafkaProducer{
		// kafkaBroker: kafkaBroker
	}
}

func (kp *KafkaProducer) PostKafka(ctx context.Context, req *pb.PostRequest, rsp *pb.PostResponse) error {
	log.Infof("Received KafkaProducer.PostRequest request: %+v", req)
	danmaku := req.Danmaku
	json_danmaku, _ := json.Marshal(danmaku)

	msg := &broker.Message{
		Header: map[string]string{
			"id": fmt.Sprintf("%d", req.ChannelID),
		},
		Body: []byte(string(json_danmaku)),
	}
	if err := broker.Publish(common.TOPIC, msg); err != nil {
		log.Errorf("Publish failed with: %+v, msg: %+v", err, msg)
		rsp.Code = 1
		rsp.Msg = err.Error()

		return nil
	} else {
		log.Infof("Publish success with msg: %+v", msg)
		return nil
	}
}
