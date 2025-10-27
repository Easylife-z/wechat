package util

// 微信返回的通用错误
type CommonError struct {
	apiName string
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
