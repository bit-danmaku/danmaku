module github.com/bit-danmaku/danmaku

go 1.17

require (
	github.com/asim/go-micro/plugins/broker/kafka/v3 v3.7.0
	github.com/asim/go-micro/plugins/server/http/v3 v3.0.0-20210924081004-8c39b1e1204d
	github.com/asim/go-micro/v3 v3.6.0
	github.com/gin-gonic/gin v1.7.4
	github.com/go-redis/redis/v8 v8.11.3
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.16
)

require github.com/gin-contrib/cors v1.3.1 // indirect
