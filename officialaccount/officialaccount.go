package officialaccount

import (
	"github.com/Easylife-z/wechat/credential"
	"github.com/Easylife-z/wechat/officialaccount/config"
	"github.com/Easylife-z/wechat/officialaccount/context"
)

// OfficialAccount 微信公众号相关API
type OfficialAccount struct {
	ctx *context.Context
}

// NewOfficialAccount 实例化公众号API
func NewOfficialAccount(cfg *config.Config) *OfficialAccount {
	var defaultAkHandle credential.AccessTokenContextHandle
	const cacheKeyPrefix = credential.CacheKeyOfficialAccountPrefix
	if cfg.UseStableAK {
		defaultAkHandle = credential.NewStableAccessToken(cfg.AppID, cfg.AppSecret, cacheKeyPrefix, cfg.Cache)
	} else {
		defaultAkHandle = credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, cacheKeyPrefix, cfg.Cache)
	}
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &OfficialAccount{ctx: ctx}
}

// SetAccessTokenHandle 自定义access_token获取方式
func (officialAccount *OfficialAccount) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	officialAccount.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (officialAccount *OfficialAccount) GetContext() *context.Context {
	return officialAccount.ctx
}
