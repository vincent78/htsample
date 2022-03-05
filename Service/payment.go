package Service

import "htSample/Model"

type IPaymentServer interface {
	PaymentList() Model.Result
	PaymentListByToken(token string) Model.Result
	PaymentListByAccount(account string) Model.Result
}

type PaymentServer struct {
}

func (p PaymentServer) PaymentList() Model.Result {
	l, e := Model.FindPaymentPOList()
	if e == nil {
		return Model.SuccessResult(l)
	} else {
		return Model.FailureResult(e.Error())
	}
}

func (p PaymentServer) PaymentListByToken(token string) Model.Result {
	l, e := Model.FindPaymentPOListByToken(token)
	if e == nil {
		return Model.SuccessResult(l)
	} else {
		return Model.FailureResult(e.Error())
	}
}

func (p PaymentServer) PaymentListByAccount(account string) Model.Result {
	l, e := Model.FindPaymentPOListByAccount(account)
	if e == nil {
		return Model.SuccessResult(l)
	} else {
		return Model.FailureResult(e.Error())
	}
}
