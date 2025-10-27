package openplatform

import (
	"github.com/Easylife-z/wechat/openplatform/config"
	"github.com/Easylife-z/wechat/openplatform/context"
	"github.com/Easylife-z/wechat/openplatform/officialaccount"
)

// OpenPlatform 微信开放平台相关api
type OpenPlatform struct {
	*context.Context
}

func NewOpenPlatform(cfg *config.Config) *OpenPlatform {
	ctx := &context.Context{
		Config: cfg,
	}
	return &OpenPlatform{ctx}
}

// GetOfficialAccount 获取要代处理的公众号（appID就是要代处理的公众号）
func (openPlatform *OpenPlatform) GetOfficialAccount(appID string) *officialaccount.OfficialAccount {
	return officialaccount.NewOfficialAccount(openPlatform.Context, appID)
}
