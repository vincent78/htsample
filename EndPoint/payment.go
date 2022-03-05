package EndPoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"htSample/Model"
	"htSample/Service"
	"htSample/global"
)

type PaymentListRequest struct {
}

type PaymentListByTokenRequest struct {
	Token string `json:"token"`
}

type PaymentListByAccountRequest struct {
	Account string `json:"account"`
}

type PaymentResponse struct {
	Model.Result
}

func MakeServerEndPointPaymentList(s Service.IPaymentServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return PaymentResponse{Result: s.PaymentList()}, nil
	}
}

func MakeServerEndPointPaymentListByToken(s Service.IPaymentServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(PaymentListByTokenRequest)
		if !ok {
			e := fmt.Errorf("the request is not PaymentRequest")
			return &PaymentResponse{Model.Result{
				Code:    global.FAILURE,
				Error:   e,
				Context: nil,
			}}, e
		} else {
			return &PaymentResponse{
				s.PaymentListByToken(r.Token),
			}, nil
		}
	}
}

func MakeServerEndPointPaymentListByAccount(s Service.IPaymentServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(PaymentListByAccountRequest)
		if !ok {
			e := fmt.Errorf("the request is not PaymentRequest")
			return &PaymentResponse{Model.Result{
				Code:    global.FAILURE,
				Error:   e,
				Context: nil,
			}}, e
		} else {
			return &PaymentResponse{
				s.PaymentListByAccount(r.Account),
			}, nil
		}
	}
}
