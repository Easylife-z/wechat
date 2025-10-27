package main

import (
	"context"
	"fmt"
	"github.com/Easylife-z/wechat/openplatform"
)

func GetRedirectUrl(opf *openplatform.OpenPlatform) (url string) {

	appID := "wx50b8f6e282e48fa5" // 第三方平台代公众号授权，我方自己的公众号-兔U片场
	scope := "snsapi_userinfo"
	redirect := "https://tuh5.mixs.cn/bind"

	platformOfficialAccount := opf.GetOfficialAccount(appID)
	oauth := platformOfficialAccount.PlatformOauth()

	url, err := oauth.GetRedirectURL(redirect, scope, "", appID)
	if err != nil {
		fmt.Printf("GetRedirectUrl err:%v", err)
	}
	return url
}

func GetUserInfo(opf *openplatform.OpenPlatform, code string) {
	appID := "wx50b8f6e282e48fa5" // 第三方平台代公众号授权，我方自己的公众号-兔U片场
	platformOfficialAccount := opf.GetOfficialAccount(appID)
	oauth := platformOfficialAccount.PlatformOauth()

	componentAccessToken, err := opf.GetComponentAccessTokenContext(context.Background())

	re, err := oauth.GetUserAccessToken(code, appID, componentAccessToken)
	fmt.Println(re, err)

	//	通过网页授权 access_token 获取用户基本信息（需授权作用域为 snsapi_userinfo）
	oaOuth := platformOfficialAccount.GetOauth()
	wxUserInfo, err := oaOuth.GetUserInfo(re.AccessToken, re.OpenID, "zh_CN")
	fmt.Println(wxUserInfo, err)
}
