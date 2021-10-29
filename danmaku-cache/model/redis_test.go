package model

import (
	"context"
	"testing"
	"reflect"
)

func TestRedisReadAndWrite(t *testing.T) {
	rdsClient := InitClient()
	ctx := context.Background()

	rdsClient.CleanChannel(ctx, 0)

	value, err := rdsClient.GetDanmakuByChannel(ctx, 0)

	if value != nil || err == nil {
		t.Errorf("get value error,value: %+v, err: %+v.", value, err)
	}
	
	dmk := []Danmaku{{ID:123, ChannelID: 0, Author: "test"}}

	err = rdsClient.SetDanmakusByChannel(ctx, 0, dmk);

	if err != nil{
		t.Errorf("failed")
	}

	value, err = rdsClient.GetDanmakuByChannel(ctx, 0)

	if value == nil || err != nil {
		t.Errorf("get value")
	}
	
	if !reflect.DeepEqual(value, dmk) {
		t.Errorf("not equal")	
	}

	err = rdsClient.CleanChannel(ctx, 0)
	if err != nil {
		t.Errorf("clean failed")
	}

	t.Cleanup(func() {
		rdsClient.CleanChannel(ctx, 0)
	})
}


