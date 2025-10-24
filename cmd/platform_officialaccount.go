package main

import (
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
