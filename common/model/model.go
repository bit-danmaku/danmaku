package model

import (
	"context"
	"time"

	log "github.com/asim/go-micro/v3/logger"
)

type Danmaku struct {
	ID        uint   `gorm:"primaryKey"`
	ChannelID uint64 `gorm:"index;not null;"`
	Author    string
	Time      float64
	Text      string
	Color     uint32
	Type      uint8     `default:"0"`
	CreatedAt time.Time `json:"-"`
}

type DBConnector struct {
	db    MariaDBClient
	cache RedisClient
}

func InitDB() *DBConnector {
	return &DBConnector{db: InitMariaDB(), cache: InitClient()}
}

// read from both cache and db.
func (db *DBConnector) GetDanmakuListByChannel(ctx context.Context, channelID uint64) []Danmaku {
	// If Hit then return cache.
	dmks, err := db.cache.GetDanmakuByChannel(ctx, channelID)
	if err != nil {
		// Or read from mysql.
		log.Info("Failed to hit Cache")

		dmks = db.db.GetDanmakuListByChannel(channelID)

		// Update cache, then return.
		err := db.cache.SetDanmakusByChannel(ctx, channelID, dmks)
		if err != nil {
			log.Error("cache set error: %+v", err)
		}
	} else {
		log.Info("Hit Cache")
	}

	return dmks
}

func (db *DBConnector) AddDanmaku(ctx context.Context, danmaku Danmaku) error {

	// Add to database.
	err := db.db.AddNewDanmaku(danmaku)

	if err != nil {
		log.Fatalf("cannot add danmaku to db, danmaku: %+v, db: %+v", danmaku, db)
	}

	// flush the cache by delete key.
	db.cache.CleanChannel(ctx, danmaku.ChannelID)

	return nil
}
