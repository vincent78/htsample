package Model

import (
	"fmt"
	"htSample/global"
)

var errorPaymentProcessNull = fmt.Errorf("paymentProcess instance is nil")

type PaymentProcessPO struct {
	Id       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Account  string `json:"account"`
	PType    string `gorm:"column:ptype" json:"ptype"` // sys_dict_type:1
	Curr     string `json:"curr"`                      // sys_dict_type:2
	Num      int64  `json:"num"`                       // 以0.01元为单位
	Token    string `json:"token"`
	Seq      int    `json:"seq"`
	Remark   string `json:"remark"`
	CreateAt int    `gorm:"autoUpdateTime:milli" json:"createAt"`
	CreateBy string `gorm:"column:create_by" json:"createBy"`
}

func NewPaymentProcessPO(a, pt, c string, n int64, tk string, s int) PaymentProcessPO {
	return PaymentProcessPO{
		Id:       0,
		Account:  a,
		PType:    pt,
		Curr:     c,
		Num:      n,
		Token:    tk,
		Seq:      s,
		Remark:   "",
		CreateAt: 0,
		CreateBy: "admin",
	}
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
