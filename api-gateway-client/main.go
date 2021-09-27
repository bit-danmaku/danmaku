package main

import (
	"context"
	"strconv"

	pb "github.com/bit-danmaku/danmaku/proto/kafkaconsumer"

	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
)

var (
	name    = "api-gateway"
	version = "latest"
)

func main() {
	httpSrv := httpServer.NewServer(
		server.Name(name),
		server.Address(":8080"),
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// register router
	demo := newDemo()
	demo.InitRouter(router)

	hd := httpSrv.NewHandler(router)
	if err := httpSrv.Handle(hd); err != nil {
		log.Fatal(err)
	}

	// Create Service
	service := micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Server(httpSrv),
		micro.Registry(registry.NewRegistry()),
	)

	service.Init()

	c := pb.NewKafkaConsumerService("kafka-consumer", service.Client())

	ret, err := c.Call(context.Background(), &pb.CallRequest{Name: "John"})

	if err != nil {
		log.Fatal(err)
	}
	log.Info(ret)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//demoRouter
type demoRouter struct {
}

type danmaku struct {
	Author string  `json:"author" binding:"required"`
	Time   float64 `json:"time" binding:"required"`
	Text   string  `json:"text" binding:"required"`
	Color  uint32  `json:"color"`
	Type   uint8   `json:"type"`
}

// [float64, uint8, uint32, string, string]
type danmakuResp = [5]interface{}

func newDemo() *demoRouter {
	return &demoRouter{}
}

func (a *demoRouter) InitRouter(router *gin.Engine) {
	router.POST("/channel/:id", a.PostDanmaku)
	router.GET("/channel/:id", a.GetDanmakuList)
}

func (a *demoRouter) PostDanmaku(c *gin.Context) {
	_, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(501, gin.H{"code": 1, "msg": "Failed When Parse Channnel ID."})
		return
	}
	strconv.Atoi(c.Param("id"))
	var dmk danmaku

	if err := c.ShouldBindJSON(&dmk); err == nil {
		log.Infof("get body: %+v", dmk)
		// TODO: call RPC client
	} else {
		log.Error(err)
		c.JSON(501, gin.H{"code": 2, "msg": "Failed When Get Post Data."})
	}

	c.JSON(200, gin.H{"code": 0, "data": dmk})
}

func (a *demoRouter) GetDanmakuList(c *gin.Context) {
	//channelID := c.Param("id")

	data := danmakuResp{1, 2, 3, "author", "hello world"}

	c.JSON(200, gin.H{"data": []danmakuResp{data, data}, "code": 0})
}
