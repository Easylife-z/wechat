package context

import (
	"github.com/Easylife-z/wechat/credential"
	"github.com/Easylife-z/wechat/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
