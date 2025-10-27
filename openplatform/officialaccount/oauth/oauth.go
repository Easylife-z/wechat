package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/Easylife-z/wechat/officialaccount/context"
	officialOauth "github.com/Easylife-z/wechat/officialaccount/oauth"
	"github.com/Easylife-z/wechat/util"
	"net/url"
)

const (
	// 授权链接
	platformRedirectOauthURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s&component_appid=%s#wechat_redirect"
	// 获取用户授权 access_token
	platformAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/component/access_token?appid=%s&code=%s&grant_type=authorization_code&component_appid=%s&component_access_token=%s"
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

/*
第三方平台 - 获取授权链接
appID：授权appID，比如兔U梦工厂申请你的昵称、头像
oauth.AppID：是我们的component_id，代授权的第三方平台appID
*/
func (oauth *Oauth) GetRedirectURL(redirectURI, scope, state, appID string) (string, error) {
	// url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(platformRedirectOauthURL, appID, urlStr, scope, state, oauth.AppID), nil
}

// GetUserAccessToken 第三方平台 - 通过网页授权的code 换取access_token(区别于context中的access_token)
func (oauth *Oauth) GetUserAccessToken(code, appID, componentAccessToken string) (result officialOauth.ResAccessToken, err error) {
	urlStr := fmt.Sprintf(platformAccessTokenURL, appID, code, oauth.AppID, componentAccessToken)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
