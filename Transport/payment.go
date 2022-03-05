package Transport

import (
	"context"
	"encoding/json"
	"errors"
	"htSample/EndPoint"
	"net/http"
)

func PaymentListDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	return EndPoint.PaymentListRequest{}, nil
}

func PaymentListEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	// 使用内置json包转换
	return json.NewEncoder(w).Encode(response)
}

func PaymentByTokenDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	p := request.URL.Query().Get("token")
	if p == "" {
		return nil, errors.New("无效参数")
	}

	return EndPoint.PaymentListByTokenRequest{Token: p}, nil
}

func PaymentByTokenEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	// 使用内置json包转换
	return json.NewEncoder(w).Encode(response)
}

func PaymentByAccountDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	p := request.URL.Query().Get("account")
	if p == "" {
		return nil, errors.New("无效参数")
	}

	return EndPoint.PaymentListByAccountRequest{Account: p}, nil
}

func PaymentByAccountEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	// 使用内置json包转换
	return json.NewEncoder(w).Encode(response)
}
