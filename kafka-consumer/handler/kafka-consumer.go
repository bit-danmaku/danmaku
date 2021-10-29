package handler

import (
	"context"

	log "github.com/asim/go-micro/v3/logger"

	pb "github.com/bit-danmaku/danmaku/proto/kafkaconsumer"
)

type KafkaConsumer struct{}

func (e *KafkaConsumer) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received KafkaConsumer.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

