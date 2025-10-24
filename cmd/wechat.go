package main

import (
	"context"
	"github.com/Easylife-z/wechat"
	"github.com/Easylife-z/wechat/cache"
)

func InitWechat() *wechat.Wechat {
	wc := wechat.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:     "127.0.0.1:6379",
		Database: 0,
	}
	ctx := context.Background()
	redisCache := cache.NewRedis(ctx, redisOpts)
	wc.SetCache(redisCache)
	return wc
}
