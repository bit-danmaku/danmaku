package handler

import (
	"context"
	"github.com/bit-danmaku/danmaku/common/model"

	log "github.com/asim/go-micro/v3/logger"
	pb "github.com/bit-danmaku/danmaku/proto/kafkaconsumer"
	_ "github.com/asim/go-micro/plugins/broker/kafka/v3"
)

type KafkaConsumer struct{}

func (e *KafkaConsumer) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received KafkaConsumer.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func sub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		//fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		
		danmaku:=Danmaku{}
    	err:=json.Unmarshal(string(p.Message().Body),&danmaku)
		if err!=nil{
			fmt.Println(err)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}

