package wechat

import (
	"github.com/Easylife-z/wechat/cache"
	"github.com/Easylife-z/wechat/openplatform"
	openConfig "github.com/Easylife-z/wechat/openplatform/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

type Wechat struct {
	cache cache.Cache // cache是个interface
}

func NewWechat() *Wechat {
	return &Wechat{}
}

func (wc *Wechat) SetCache(cache cache.Cache) {
	wc.cache = cache
}

// GetOpenPlatform 获取微信开放平台的实例
func (wc *Wechat) GetOpenPlatform(cfg *openConfig.Config) *openplatform.OpenPlatform {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return openplatform.NewOpenPlatform(cfg)
}
