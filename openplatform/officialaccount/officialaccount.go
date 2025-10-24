package officialaccount

import (
	"github.com/Easylife-z/wechat/credential"
	"github.com/Easylife-z/wechat/officialaccount"
	offConfig "github.com/Easylife-z/wechat/officialaccount/config"
	opContext "github.com/Easylife-z/wechat/openplatform/context"
	"github.com/Easylife-z/wechat/openplatform/officialaccount/oauth"
)

// OfficialAccount 代公众号实现业务
type OfficialAccount struct {
	// 授权的公众号的appID
	appID                            string
	*officialaccount.OfficialAccount // 代授权的那个（第三方平台-蜜阅FM）
}

// NewOfficialAccount 实例化
// appID :为授权方公众号 APPID，非开放平台第三方平台 APPID
func NewOfficialAccount(opCtx *opContext.Context, appID string) *OfficialAccount {
	officialAccount := officialaccount.NewOfficialAccount(&offConfig.Config{
		AppID:          opCtx.AppID,
		EncodingAESKey: opCtx.EncodingAESKey,
		Token:          opCtx.Token,
		Cache:          opCtx.Cache,
	})
	// 设置获取access_token的函数
	officialAccount.SetAccessTokenHandle(NewDefaultAuthrAccessToken(opCtx, appID))
	return &OfficialAccount{appID: appID, OfficialAccount: officialAccount}
}

// PlatformOauth 平台代发起oauth2网页授权
func (officialAccount *OfficialAccount) PlatformOauth() *oauth.Oauth {
	return oauth.NewOauth(officialAccount.GetContext())
}

// DefaultAuthrAccessToken 默认获取授权ak的方法
type DefaultAuthrAccessToken struct {
	opCtx *opContext.Context
	appID string
}

// NewDefaultAuthrAccessToken New
func NewDefaultAuthrAccessToken(opCtx *opContext.Context, appID string) credential.AccessTokenHandle {
	return &DefaultAuthrAccessToken{
		opCtx: opCtx,
		appID: appID,
	}
}

// GetAccessToken 获取ak
func (ak *DefaultAuthrAccessToken) GetAccessToken() (string, error) {
	return ak.opCtx.GetAuthrAccessToken(ak.appID)
}
