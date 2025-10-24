package config

import "github.com/Easylife-z/wechat/cache"

// Config .config for 微信开放平台
type Config struct {
	AppID          string `json:"app_id"`
	AppSecret      string `json:"app_secret"`
	Token          string `json:"token"`
	EncodingAESKey string `json:"encoding_aes_key"`
	Cache          cache.Cache
}
