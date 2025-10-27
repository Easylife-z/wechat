// Package context 开放平台相关context
package context

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Easylife-z/wechat/cache"
	"github.com/Easylife-z/wechat/util"
	"time"
)

const (
	// 获取授权账号调用令牌
	refreshTokenURL = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=%s"
)

// 获取授权账号调用令牌返回结果-授权方AccessToken
type AuthrAccessToken struct {
	Appid        string `json:"authorizer_appid"`
	AccessToken  string `json:"authorizer_access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"authorizer_refresh_token"`
}

// GetAuthrAccessToken 获取授权方AccessToken
// 会依次调用GetAuthrAccessTokenContext、RefreshAuthrTokenContext、GetComponentAccessTokenContext
func (ctx *Context) GetAuthrAccessToken(appid string) (string, error) {
	return ctx.GetAuthrAccessTokenContext(context.Background(), appid)
}

// GetAuthrAccessTokenContext 获取授权方AccessToken
func (ctx *Context) GetAuthrAccessTokenContext(stdCtx context.Context, appid string) (string, error) {
	authrTokenKey := "authorizer_access_token_" + appid
	val := cache.GetContext(stdCtx, ctx.Cache, authrTokenKey)
	if val == nil {
		refreshTokenKey := "authorizer_refresh_token_" + appid
		val := cache.GetContext(stdCtx, ctx.Cache, refreshTokenKey)
		if val == nil {
			return "", fmt.Errorf("cannot get authorizer %s refresh token", appid)
		}
		token, err := ctx.RefreshAuthrTokenContext(stdCtx, appid, val.(string))
		if err != nil {
			return "", err
		}
		return token.AccessToken, nil
	}

	return val.(string), nil
}

// RefreshAuthrTokenContext 获取（刷新）授权公众号或小程序的接口调用凭据（令牌）
func (ctx *Context) RefreshAuthrTokenContext(stdCtx context.Context, appid, refreshToken string) (*AuthrAccessToken, error) {
	cat, err := ctx.GetComponentAccessTokenContext(stdCtx)
	if err != nil {
		return nil, err
	}

	req := map[string]string{
		"component_appid":          ctx.AppID,
		"authorizer_appid":         appid,
		"authorizer_refresh_token": refreshToken,
	}
	uri := fmt.Sprintf(refreshTokenURL, cat)
	body, err := util.PostJSONContext(stdCtx, uri, req)
	if err != nil {
		return nil, err
	}

	ret := &AuthrAccessToken{}
	if err := json.Unmarshal(body, ret); err != nil {
		return nil, err
	}

	authrTokenKey := "authorizer_access_token_" + appid
	if err := cache.SetContext(stdCtx, ctx.Cache, authrTokenKey, ret.AccessToken, time.Second*time.Duration(ret.ExpiresIn-30)); err != nil {
		return nil, err
	}
	refreshTokenKey := "authorizer_refresh_token_" + appid
	if err := cache.SetContext(stdCtx, ctx.Cache, refreshTokenKey, ret.RefreshToken, 10*365*24*60*60*time.Second); err != nil {
		return nil, err
	}
	return ret, nil
}

// GetComponentAccessTokenContext 获取 ComponentAccessToken
func (ctx *Context) GetComponentAccessTokenContext(stdCtx context.Context) (string, error) {
	accessTokenCacheKey := fmt.Sprintf("component_access_token_%s", ctx.AppID)
	val := cache.GetContext(stdCtx, ctx.Cache, accessTokenCacheKey)
	if val == nil {
		return "", fmt.Errorf("cann't get component access token")
	}
	return val.(string), nil
}
