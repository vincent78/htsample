package Transport

import (
	"context"
	"encoding/json"
	"errors"
	"htSample/EndPoint"
	"net/http"
)

func AccountByCodeDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	code := request.URL.Query().Get("code")
	if code == "" {
		return nil, errors.New("无效参数")
	}

	return EndPoint.AccountRequest{Code: code}, nil
}

func AccountByCodeEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	// 使用内置json包转换
	return json.NewEncoder(w).Encode(response)
}

func AccountListDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	return struct {
	}{}, nil
}

func AccountListEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	// 使用内置json包转换
	return json.NewEncoder(w).Encode(response)
}

func AccountBalanceByCodeDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	code := request.URL.Query().Get("code")
	if code == "" {
		return nil, errors.New("无效参数")
	}

	return EndPoint.AccountRequest{Code: code}, nil
}

func AccountBalanceByCodeEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	// 使用内置json包转换
	return json.NewEncoder(w).Encode(response)
}
