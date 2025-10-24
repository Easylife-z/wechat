package util

// CommonError 微信返回的通用错误 json
type CommonError struct {
	apiName string
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
