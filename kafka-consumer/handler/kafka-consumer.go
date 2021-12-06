package handler

import (
	"context"
	"github.com/bit-danmaku/danmaku/common/model"

	log "github.com/asim/go-micro/v3/logger"
	pb "github.com/bit-danmaku/danmaku/proto/kafkaconsumer"
)

type KafkaConsumer struct{
	dbConnector *model.DBConnector
}


func (kc *KafkaConsumer) PostDanmaku(ctx context.Context, req *pb.PostRequest, rsp *pb.PostResponse) error {
	log.Infof("Received DanmakuCache.GetDanmakuListByChannel request: %+v", req)

	danmaku := req.Danmaku

	err := kc.dbConnector.AddDanmaku(ctx, model.Danmaku{
		ChannelID: req.ChannelID,
		Author:    danmaku.Author,
		Time:      danmaku.Time,
		Text:      danmaku.Text,
		Color:     danmaku.Color,
		Type:      uint8(danmaku.Type),
	})

	if err != nil {
		return err
	}

	return nil
}

