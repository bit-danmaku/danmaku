package handler

import (
	"context"

	log "github.com/asim/go-micro/v3/logger"

	"github.com/bit-danmaku/danmaku/common/model"
	"github.com/bit-danmaku/danmaku/proto/common"
	pb "github.com/bit-danmaku/danmaku/proto/danmakucache"
)

type DanmakuCache struct {
	dbConnector *model.DBConnector
}

func InitDanmakuCache() DanmakuCache {
	return DanmakuCache{
		dbConnector: model.InitDB(),
	}
}

func (dc *DanmakuCache) GetDanmakuListByChannel(ctx context.Context, req *pb.GetRequest, rsp *pb.GetResponse) error {
	log.Infof("Received DanmakuCache.GetDanmakuListByChannel request: %+v", req)
	danmakuList := dc.dbConnector.GetDanmakuListByChannel(ctx, req.ChannelID)

	for _, v := range danmakuList {
		rsp.DanmakuList = append(rsp.DanmakuList, &common.Danmaku{
			Author: v.Author,
			Time:   v.Time,
			Text:   v.Text,
			Color:  v.Color,
			Type:   uint32(v.Type),
		})
	}
	return nil
}

