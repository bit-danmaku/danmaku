package handler

import (
	"context"
	"fmt"
	"encoding/json"
	"github.com/asim/go-micro/v3/broker"
	"github.com/bit-danmaku/danmaku/common/model"

	_ "github.com/asim/go-micro/plugins/broker/kafka/v3"
	log "github.com/asim/go-micro/v3/logger"
	commonProto "github.com/bit-danmaku/danmaku/proto/common"
	pb "github.com/bit-danmaku/danmaku/proto/kafkaconsumer"
)

var (
	service = "kafka-consumer"
	version = "latest"
)
var (
	topic = "go.micro.topic.foo"
)

type KafkaConsumer struct{
	dbConnector *model.DBConnector
}

func (kc *KafkaConsumer) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received KafkaConsumer.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (kc *KafkaConsumer) sub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		//fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)

		danmaku:=commonProto.Danmaku{}
		if err:=json.Unmarshal(p.Message().Body,&danmaku);err!=nil{
			fmt.Println(err)
		}

		err := kc.dbConnector.AddDanmaku(ctx, model.Danmaku{
			ChannelID: req.ChannelID,
			Author:    danmaku.Author,
			Time:      danmaku.Time,
			Text:      danmaku.Text,
			Color:     danmaku.Color,
			Type:      uint8(danmaku.Type),
		})

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}



//func (kc *KafkaConsumer) PostDanmaku(ctx context.Context, req *pb.PostRequest, rsp *pb.PostResponse) error {
//	log.Infof("Received DanmakuCache.GetDanmakuListByChannel request: %+v", req)
//
//	danmaku := req.Danmaku
//
//	err := kc.dbConnector.AddDanmaku(ctx, model.Danmaku{
//		ChannelID: req.ChannelID,
//		Author:    danmaku.Author,
//		Time:      danmaku.Time,
//		Text:      danmaku.Text,
//		Color:     danmaku.Color,
//		Type:      uint8(danmaku.Type),
//	})
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

