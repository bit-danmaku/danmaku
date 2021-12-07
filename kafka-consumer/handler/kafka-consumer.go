package handler

import (
	"context"
	"encoding/json"
	"github.com/asim/go-micro/v3/broker"
	"github.com/bit-danmaku/danmaku/common/model"
	"strconv"

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

func InitKafkaConsumer() KafkaConsumer {
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	return KafkaConsumer{}

}

func (kc *KafkaConsumer) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received KafkaConsumer.Call request: %v", req)

	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		//fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)

		danmaku:=commonProto.Danmaku{}
		if err_json:=json.Unmarshal(p.Message().Body,&danmaku);err_json!=nil{
			return err_json
		}

		channelID, err_int := strconv.ParseUint(p.Message().Header["id"], 10, 0)
		if err_int != nil {
			return err_int
		}

		err_addDanmaku := kc.dbConnector.AddDanmaku(ctx, model.Danmaku{
			ChannelID: channelID,
			Author:    danmaku.Author,
			Time:      danmaku.Time,
			Text:      danmaku.Text,
			Color:     danmaku.Color,
			Type:      uint8(danmaku.Type),
		})
		if err_addDanmaku != nil{
			return err_addDanmaku
		}
		return nil
	})

	if err != nil {
		rsp.Code = 1
		rsp.Msg = err.Error()
		return err
	}else {
		return nil
	}

}

//func (kc *KafkaConsumer) sub() {
//	_, err := broker.Subscribe(topic, func(p broker.Event) error {
//		//fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
//
//		danmaku:=commonProto.Danmaku{}
//		if err:=json.Unmarshal(p.Message().Body,&danmaku);err!=nil{
//			fmt.Println(err)
//		}
//
//		err := kc.dbConnector.AddDanmaku(ctx, model.Danmaku{
//			ChannelID: req.ChannelID,
//			Author:    danmaku.Author,
//			Time:      danmaku.Time,
//			Text:      danmaku.Text,
//			Color:     danmaku.Color,
//			Type:      uint8(danmaku.Type),
//		})
//
//		return nil
//	})
//	if err != nil {
//		fmt.Println(err)
//	}
//
//}



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

