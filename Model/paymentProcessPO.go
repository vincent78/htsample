package Model

import (
	"fmt"
	"htSample/global"
)

var errorPaymentProcessNull = fmt.Errorf("paymentProcess instance is nil")

type PaymentProcessPO struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Account  string ``
	PType    string `gorm:"column:ptype"` // sys_dict_type:1
	Curr     string ``                    // sys_dict_type:2
	Balance  int    `gorm:"balance"`      // 以0.01元为单位
	Frozen   int    `gorm:"frozen"`       // 以0.01元为单位
	Token    string ``
	Seq      int    `gorm:"seq"`
	Remark   string `gorm:"remark"`
	CreateAt int    `gorm:"autoUpdateTime:milli"`
	CreateBy string `gorm:"column:create_by"`
}

func (a PaymentProcessPO) TableName() string {
	return "payment_process"
}

func FindPaymentProcessList() ([]PaymentProcessPO, error) {
	l := make([]PaymentProcessPO, 0)
	r := global.DB.Order("token desc, seq asc").Find(&l)
	return l, r.Error
}

func FindPaymentProcessListByToken(tk string) ([]PaymentProcessPO, error) {
	l := make([]PaymentProcessPO, 0)
	r := global.DB.Where("token = ?", tk).Order("seq asc").Find(&l)
	return l, r.Error
}

func (a *PaymentProcessPO) Create() error {
	if a == nil {
		return errorPaymentProcessNull
	}
	r := global.DB.Create(a)
	return r.Error
}
