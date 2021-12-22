package main

import (
	"context"
	"strconv"

	danmaku_cache_pb "github.com/bit-danmaku/danmaku/proto/danmakucache"
	kafka_producer_pb "github.com/bit-danmaku/danmaku/proto/kafkaproducer"

	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/bit-danmaku/danmaku/common"
	commonProto "github.com/bit-danmaku/danmaku/proto/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	version = "latest"
)

func main() {
	httpSrv := httpServer.NewServer(
		server.Name(common.API_GATEWAY),
		server.Address(":8080"),
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// register router
	hd := httpSrv.NewHandler(router)
	if err := httpSrv.Handle(hd); err != nil {
		log.Fatal(err)
	}

	// Create Service
	service := micro.NewService(
		micro.Name(common.API_GATEWAY),
		micro.Version(version),
		micro.Server(httpSrv),
		micro.Registry(registry.NewRegistry()),
	)

	service.Init()

	demo := newDemo(service.Client())
	demo.InitRouter(router)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//demoRouter
type demoRouter struct {
	danmakuCachePB danmaku_cache_pb.DanmakuCacheService
	kafkaProducer  kafka_producer_pb.KafkaProducerService
}

type danmaku struct {
	Author string  `json:"author" binding:"required"`
	Time   float64 `json:"time"`
	Text   string  `json:"text" binding:"required"`
	Color  uint32  `json:"color"`
	Type   uint8   `json:"type"`
}

// [float64, uint8, uint32, string, string]
type danmakuResp = [5]interface{}

func newDemo(client client.Client) *demoRouter {
	return &demoRouter{
		danmakuCachePB: danmaku_cache_pb.NewDanmakuCacheService(common.DANMAKU_CACHE, client),
		kafkaProducer:  kafka_producer_pb.NewKafkaProducerService(common.KAFKA_PRODUCER, client),
	}
}

func (a *demoRouter) InitRouter(router *gin.Engine) {
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	router.POST("/channel/:id/v3/", a.PostDanmaku)
	router.GET("/channel/:id/v3/", a.GetDanmakuList)
}

func (a *demoRouter) PostDanmaku(c *gin.Context) {
	channelID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(501, gin.H{"code": 1, "msg": "Failed When Parse Channnel ID."})
		return
	}
	var dmk danmaku

	if err := c.ShouldBindJSON(&dmk); err == nil {
		log.Infof("get body: %+v", dmk)

		ret, err := a.kafkaProducer.PostKafka(context.Background(), &kafka_producer_pb.PostRequest{Danmaku: &commonProto.Danmaku{Author: dmk.Author, Time: dmk.Time, Text: dmk.Text, Color: dmk.Color, Type: uint32(dmk.Type)}, ChannelID: channelID})

		if err != nil {
			c.JSON(501, gin.H{"code": 2, "msg": "Failed When Add Data to DB."})
			return
		}

		c.JSON(200, gin.H{"code": ret.Code, "data": ret.Msg})
		return

	} else {
		log.Error(err)
		c.JSON(501, gin.H{"code": 2, "msg": "Failed When Get Post Data."})
	}

}

func (a *demoRouter) GetDanmakuList(c *gin.Context) {
	channelID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(501, gin.H{"code": 1, "msg": "Failed When Parse Channnel ID."})
		return
	}

	ret, err := a.danmakuCachePB.GetDanmakuListByChannel(context.Background(), &danmaku_cache_pb.GetRequest{ChannelID: channelID})

	danmankuList := make([]danmakuResp, len(ret.DanmakuList))

	if err != nil {
		log.Error(err)
		c.JSON(501, gin.H{"data": danmankuList, "code": 2, "msg": "Failed When Get Data."})
		return
	}

	for i, v := range ret.DanmakuList {
		danmankuList[i] = danmakuResp{v.Time, v.Type, v.Color, v.Author, v.Text}
	}

	c.JSON(200, gin.H{"data": danmankuList, "code": 0})
}
