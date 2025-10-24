package config

import "github.com/Easylife-z/wechat/cache"

// Config .config for 微信公众号
type Config struct {
	AppID          string `json:"app_id"`           // appid
	AppSecret      string `json:"app_secret"`       // appsecret
	Token          string `json:"token"`            // token
	EncodingAESKey string `json:"encoding_aes_key"` // EncodingAESKey
	Cache          cache.Cache
	UseStableAK    bool // use the stable access_token
}
