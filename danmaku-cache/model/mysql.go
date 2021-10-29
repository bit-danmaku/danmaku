package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MariaDBClient struct {
	db *gorm.DB
}

func InitMariaDB() MariaDBClient {
	dsn := "root:qwerty@tcp(db.danmaku.fkynjyq.com:3306)/olddanmaku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate table structure.
	db.AutoMigrate(&Danmaku{})

	return MariaDBClient{db}
}

func (db *MariaDBClient) GetDanmakuListByChannel(channelID uint64) []Danmaku {
	var danmakus []Danmaku
	db.db.Where(&Danmaku{ChannelID: channelID}).Find(&danmakus)

	return danmakus
}
