package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// URIModifier URI修改器
type URIModifier func(uri string) string

var uriModifier URIModifier

// DefaultHTTPClient 默认httpClient
var DefaultHTTPClient = http.DefaultClient

// HTTPGetContext get 请求
func HTTPGetContext(ctx context.Context, uri string) ([]byte, error) {
	if uriModifier != nil {
		uri = uriModifier(uri)
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	response, err := DefaultHTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return io.ReadAll(response.Body)
}

// PostJSONContext post json 数据请求
func PostJSONContext(ctx context.Context, uri string, obj interface{}) ([]byte, error) {
	if uriModifier != nil {
		uri = uriModifier(uri)
	}
	jsonBuf := new(bytes.Buffer)
	enc := json.NewEncoder(jsonBuf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(obj)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", uri, jsonBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	response, err := DefaultHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return io.ReadAll(response.Body)
}
