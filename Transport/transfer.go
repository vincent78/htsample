package Transport

import (
	"context"
	"encoding/json"
	"htSample/EndPoint"
	"net/http"
	"strconv"
)

func TransferDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	// 根据请求body创建一个json解析器实例
	decoder := json.NewDecoder(request.Body)
	// 用于存放参数key=value数据
	var params map[string]string
	// 解析参数 存入map
	decoder.Decode(&params)
	n1, e := strconv.ParseInt(params["num1"], 10, 64)
	if e != nil {
		return nil, e
	}
	n2, e := strconv.ParseInt(params["num2"], 10, 64)
	if e != nil {
		return nil, e
	}
	return EndPoint.TransferRequest{
		Account1: params["account1"],
		Curr1:    params["curr1"],
		Num1:     n1,
		Account2: params["account2"],
		Curr2:    params["curr2"],
		Num2:     n2,
	}, nil
}

func TransferByCodeEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	// 使用内置json包转换
	return json.NewEncoder(w).Encode(response)
}
