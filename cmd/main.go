package main

import "fmt"

func main() {
	// 1、初始化wechat
	wc := InitWechat()

	//	2、获取微信开放平台的实例
	opf := InitOpenPlatform(wc)

	// 3、公众号代处理

	// 3.1 获取授权链接
	oauthUrl := GetRedirectUrl(opf)
	fmt.Println(oauthUrl)

	//	3.2 获取微信用户基本信息
	GetUserInfo(opf, "")
}
