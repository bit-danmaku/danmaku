package handler

import (
	"context"
	"encoding/json"
	"fmt"
	kafka "github.com/asim/go-micro/plugins/broker/kafka/v3"
	"github.com/asim/go-micro/v3/broker"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/bit-danmaku/danmaku/common/model"
	commonProto "github.com/bit-danmaku/danmaku/proto/common"
	pb "github.com/bit-danmaku/danmaku/proto/kafkaconsumer"
	"strconv"
)

var (
	service = "kafka-consumer"
	version = "latest"
)
var (
	topic = "danmaku"
)



type KafkaConsumer struct{
	dbConnector *model.DBConnector
	kafkaBroker broker.Broker
}

func InitKafkaConsumer() *KafkaConsumer {
	kafkaBroker := kafka.NewBroker()
	if err := kafkaBroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := kafkaBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	return &KafkaConsumer{dbConnector: model.InitDB(), kafkaBroker: kafkaBroker}
}

func (kc *KafkaConsumer) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received KafkaConsumer.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (kc *KafkaConsumer) Sub() {
	_, err := kc.kafkaBroker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("收到订阅")
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)

		danmaku:=commonProto.Danmaku{}
		if err_json:=json.Unmarshal(p.Message().Body, &danmaku);err_json!=nil{
			fmt.Println(err_json)
			return err_json
		}

		channelID, err_int := strconv.ParseUint(p.Message().Header["id"], 10, 0)
		if err_int != nil {
			return err_int
		}

		fmt.Println("解析完成")

		
		err_addDanmaku := kc.dbConnector.AddDanmaku(context.Background(), model.Danmaku{
			ChannelID: channelID,
			Author:    danmaku.Author,
			Time:      danmaku.Time,
			Text:      danmaku.Text,
			Color:     danmaku.Color,
			Type:      uint8(danmaku.Type),
		})
		if err_addDanmaku != nil{
			fmt.Println("err_addDanmaku")
			return err_addDanmaku
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

