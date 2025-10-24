package oauth

import (
	"fmt"
	"github.com/Easylife-z/wechat/officialaccount/context"
	"net/url"
)

const (
	platformRedirectOauthURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s&component_appid=%s#wechat_redirect"
)

// Oauth 平台代发起oauth2网页授权
type Oauth struct {
	*context.Context
}

func NewOauth(context *context.Context) *Oauth {
	auth := new(Oauth)
	auth.Context = context
	return auth
}

// GetRedirectURL 第三方平台 - 获取跳转的url地址
func (oauth *Oauth) GetRedirectURL(redirectURI, scope, state, appID string) (string, error) {
	// url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(platformRedirectOauthURL, appID, urlStr, scope, state, oauth.AppID), nil
}
