package EndPoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"htSample/Model"
	"htSample/Service"
	"htSample/global"
)

type TransferRequest struct {
	Account1 string `json:"account1"`
	Curr1    string `json:"curr1"`
	Num1     int64  `json:"num1"`
	Account2 string `json:"account2"`
	Curr2    string `json:"curr2"`
	Num2     int64  `json:"num2"`
}

type TransferResponse struct {
	Model.Result
}

func MakeServerEndPointTransfer(s Service.ITransferServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(TransferRequest)
		if !ok {
			e := fmt.Errorf("the request is not TransferRequest")
			return &AccountResponse{Model.Result{
				Code:    global.FAILURE,
				Error:   e,
				Context: nil,
			}}, e
		} else {
			return TransferResponse{Result: s.Transfer(r.Account1, r.Curr1, r.Num1, r.Account2, r.Curr2, r.Num2)}, nil
		}
	}
}
