package main

import (
	"github.com/bit-danmaku/danmaku/common"
	"github.com/bit-danmaku/danmaku/danmaku-cache/handler"
	pb "github.com/bit-danmaku/danmaku/proto/danmakucache"

	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
)

var (
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(common.DANMAKU_CACHE),
		micro.Version(version),
	)
	srv.Init()

	hdlStruct := handler.InitDanmakuCache()

	// Register handler
	pb.RegisterDanmakuCacheHandler(srv.Server(), &hdlStruct)

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
