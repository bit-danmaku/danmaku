package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/asim/go-micro/plugins/broker/kafka/v3"
	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/v3/cmd"
	log "github.com/asim/go-micro/v3/logger"
	common "github.com/bit-danmaku/danmaku/common"
	"github.com/bit-danmaku/danmaku/common/model"
	commonProto "github.com/bit-danmaku/danmaku/proto/common"
	pb "github.com/bit-danmaku/danmaku/proto/kafkaconsumer"
)

var (
	version = "latest"
)

type KafkaConsumer struct {
	dbConnector *model.DBConnector
	//kafkaBroker broker.Broker
}

func InitKafkaConsumer() *KafkaConsumer {
	cmd.Init()
	//kafkaBroker := kafka.NewBroker()
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	return &KafkaConsumer{dbConnector: model.InitDB() /*kafkaBroker: kafkaBroker*/}
}

func (kc *KafkaConsumer) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received KafkaConsumer.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (kc *KafkaConsumer) Sub() {
	_, err := broker.Subscribe(common.TOPIC, func(p broker.Event) error {
		log.Info("收到订阅")
		log.Infof("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)

		danmaku := commonProto.Danmaku{}
		if err_json := json.Unmarshal(p.Message().Body, &danmaku); err_json != nil {
			fmt.Println(err_json)
			return err_json
		}

		channelID, err_int := strconv.ParseUint(p.Message().Header["id"], 10, 0)
		if err_int != nil {
			return err_int
		}

		log.Info("解析完成")

		err_addDanmaku := kc.dbConnector.AddDanmaku(context.Background(), model.Danmaku{
			ChannelID: channelID,
			Author:    danmaku.Author,
			Time:      danmaku.Time,
			Text:      danmaku.Text,
			Color:     danmaku.Color,
			Type:      uint8(danmaku.Type),
		})
		if err_addDanmaku != nil {
			log.Error("err_addDanmaku")
			return err_addDanmaku
		}
		return nil
	})

	if err != nil {
		log.Error(err)
	}
}
