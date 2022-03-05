package EndPoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"htSample/Model"
	"htSample/Service"
	"htSample/global"
)

type AccountRequest struct {
	Code string `json:"code"`
}

type AccountResponse struct {
	Model.Result
}

type AccountAllRequest struct{}

func MakeServerEndPointAccountByCode(s Service.IAccountServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(AccountRequest)
		if !ok {
			e := fmt.Errorf("the request is not AccountRequest")
			return &AccountResponse{Model.Result{
				Code:    global.FAILURE,
				Error:   e,
				Context: nil,
			}}, e
		} else {
			return AccountResponse{Result: s.AccountByCode(r.Code)}, nil
		}
	}
}

func MakeServerEndPointAccountList(s Service.IAccountServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return AccountResponse{Result: s.AccountList()}, nil
	}
}

func MakeServerEndPointAccountBalanceByCode(s Service.IAccountServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(AccountRequest)
		if !ok {
			e := fmt.Errorf("the request is not AccountRequest")
			return &AccountResponse{Model.Result{
				Code:    global.FAILURE,
				Error:   e,
				Context: nil,
			}}, e
		} else {
			return AccountResponse{Result: s.AccountBalanceByCode(r.Code)}, nil
		}
	}
}
