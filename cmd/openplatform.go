package main

import (
	"github.com/Easylife-z/wechat"
	"github.com/Easylife-z/wechat/openplatform"
	openConfig "github.com/Easylife-z/wechat/openplatform/config"
)

func InitOpenPlatform(wc *wechat.Wechat) *openplatform.OpenPlatform {
	openCfg := openConfig.Config{
		AppID:          "wxa638f9f4b227b09b",
		AppSecret:      "637bea7cc3bfd65d72a659b4761baef4",
		Token:          "NpVsaIxPkXQW5+uDLgdjBjMVcKkFjlUFQhEnNJKxkaI",
		EncodingAESKey: "NpVsaIxPkXQW5uDLgdjBjMVcKkFjlUFQhEnNJKxkaI2",
	}
	opf := wc.GetOpenPlatform(&openCfg)
	return opf
}
