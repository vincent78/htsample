package Service

import "htSample/Model"

type IAccountServer interface {
	AccountByCode(code string) Model.Result
	AccountList() Model.Result
}

type AccountServer struct {
}

func (a AccountServer) AccountByCode(c string) Model.Result {
	apo, e := Model.FindAccountByCode(c)
	if e == nil {
		return Model.SuccessResult(*apo)
	} else {
		return Model.FailureResult(e.Error())
	}
}

func (a AccountServer) AccountList() Model.Result {
	l, e := Model.FindAccountAll()
	if e == nil {
		return Model.SuccessResult(l)
	} else {
		return Model.FailureResult(e.Error())
	}
}
