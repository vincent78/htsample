package Service

import (
	"htSample/Model"
	"htSample/utils"
)

type ITransferServer interface {
	Transfer(u1 string, c1 string, n1 int64, u2 string, c2 string, n2 int64) Model.Result
}

type TransferServer struct {
}

func (t TransferServer) Transfer(u1 string, c1 string, n1 int64, u2 string, c2 string, n2 int64) Model.Result {
	// check u1
	uo1, e := Model.FindAccountByCode(u1)
	if e != nil {
		return Model.FailureResult("find user[%v] error: %v", u1, e.Error())
	}
	if uo1 == nil {
		return Model.FailureResult("user[%v] is not exist", u1)
	}

	ubo1, e := Model.FindAccountBalancePOByCode(u1)
	if e != nil {
		return Model.FailureResult("find user[%v] error: %v", u1, e.Error())
	}
	if ubo1 == nil {
		return Model.FailureResult("the user[%v] is not exist", u1)
	}

	// checking outgoing
	if ubo1.Balance < n1 {
		return Model.FailureResult("user[%v] have not enough curr[%v] balance[%v], require %v", u1, c1, ubo1.Balance, n1)
	}

	// check u2

	uo2, e := Model.FindAccountByCode(u2)
	if e != nil {
		return Model.FailureResult("find user[%v] error: %v", u2, e.Error())
	}
	if uo2 == nil {
		return Model.FailureResult("user[%v] is not exist", u2)
	}
	ubo2, e := Model.FindAccountBalancePOByCode(u2)
	if e != nil {
		return Model.FailureResult("find user[%v] error: %v", u2, e.Error())
	}
	if ubo2 == nil {
		ubo2 = &Model.AccountBalancePO{
			Id:      0,
			Code:    u2,
			Curr:    c2,
			Balance: 0,
		}
		ubo2.Save()
	}

	tk := utils.GetUUID()
	step := 1

	p1 := Model.NewPaymentPO(u1, "outgoing", c1, n1, 0, tk, "admin")
	p1.Create()

	p1.FrozenNum(n1)
	pp1 := Model.NewPaymentProcessPO(u1, "frozen", c1, n1, tk, step)
	pp1.Create()

	ubo1.Balance -= n1
	ubo1.Save()

	p2 := &Model.PaymentPO{
		Id:       0,
		Account:  u2,
		PType:    "incoming",
		Curr:     c2,
		Balance:  0,
		Frozen:   0,
		Token:    tk,
		Remark:   "",
		CreateAt: 0,
		CreateBy: "admin",
	}
	p2.Balance += n2
	p2.Create()
	step += 1
	pp2 := Model.NewPaymentProcessPO(u2, "incoming", c2, n2, tk, step)
	pp2.Create()

	ubo2.Balance += n2
	ubo2.Save()

	p1.CommitNum(n2)

	step += 1
	pp3 := Model.NewPaymentProcessPO(u1, "commit", c1, n1, tk, step)
	pp3.Create()

	return Model.SuccessResult(tk)
}
