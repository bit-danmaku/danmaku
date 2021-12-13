package model

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	rdb *redis.Client
}

func InitClient() RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		//Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return RedisClient{
		rdb,
	}
}

func (rds *RedisClient) GetDanmakuByChannel(ctx context.Context, channelID uint64) ([]Danmaku, error) {
	val, err := rds.rdb.Get(ctx, strconv.FormatUint(channelID, 10)).Result()
	if err != nil {
		// Error happens.

		return nil, err
	}

	var danmakus []Danmaku
	err = json.Unmarshal([]byte(val), &danmakus)

	if err != nil {
		return nil, err
	}
	return danmakus, nil
}

func (rds *RedisClient) SetDanmakusByChannel(ctx context.Context, channelID uint64, danmakus []Danmaku) error {
	data, err := json.Marshal(danmakus)
	if err != nil {

		return err
	}

	err = rds.rdb.Set(ctx, strconv.FormatUint(channelID, 10), data, 0).Err()

	if err != nil {

		return err
	}

	return nil
}

func (rds *RedisClient) CleanChannel(ctx context.Context, channelID uint64) error {
	err := rds.rdb.Del(ctx, strconv.FormatUint(channelID, 10)).Err()

	if err != nil {
		return err
	}

	return nil
}
