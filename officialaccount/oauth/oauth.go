package oauth

import (
	ctx2 "context"
	"encoding/json"
	"fmt"
	"github.com/Easylife-z/wechat/officialaccount/context"
	"github.com/Easylife-z/wechat/util"
)

const (
	userInfoURL = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s"
)

// Oauth 保存用户授权信息
type Oauth struct {
	*context.Context
}

// NewOauth 实例化授权信息
func NewOauth(context *context.Context) *Oauth {
	auth := new(Oauth)
	auth.Context = context
	return auth
}

// ResAccessToken 获取用户授权access_token的返回结果
type ResAccessToken struct {
	util.CommonError

	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`

	// IsSnapShotUser 是否为快照页模式虚拟账号，只有当用户是快照页模式虚拟账号时返回，值为1
	// 公众号文档 https://developers.weixin.qq.com/community/minihome/doc/000c2c34068880629ced91a2f56001
	IsSnapShotUser int `json:"is_snapshotuser"`

	// UnionID 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	// 公众号文档 https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842
	UnionID string `json:"unionid"`
}

// UserInfo 用户授权获取到用户信息
type UserInfo struct {
	util.CommonError

	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int32    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

// GetUserInfo 如果scope为 snsapi_userinfo 则可以通过此方法获取到用户基本信息
func (oauth *Oauth) GetUserInfo(accessToken, openID, lang string) (result UserInfo, err error) {
	return oauth.GetUserInfoContext(ctx2.Background(), accessToken, openID, lang)
}

// GetUserInfoContext 如果scope为 snsapi_userinfo 则可以通过此方法获取到用户基本信息 with context
func (oauth *Oauth) GetUserInfoContext(ctx ctx2.Context, accessToken, openID, lang string) (result UserInfo, err error) {
	if lang == "" {
		lang = "zh_CN"
	}
	urlStr := fmt.Sprintf(userInfoURL, accessToken, openID, lang)
	var response []byte
	response, err = util.HTTPGetContext(ctx, urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
