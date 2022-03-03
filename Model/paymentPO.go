package Model

import (
	"fmt"
	"htSample/global"
)

type PaymentPO struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Account  string ``
	PType    string `gorm:"column:ptype"` // sys_dict_type:1
	Curr     string ``                    // sys_dict_type:2
	Balance  int    ``                    // 以0.01元为单位
	Frozen   int    ``                    // 以0.01元为单位
	Token    string ``
	Remark   string `gorm:"remark"`
	CreateAt int    `gorm:"autoUpdateTime:milli"`
	CreateBy string `gorm:"column:create_by"`
}

func (p PaymentPO) TableName() string {
	return "payment"
}

func NewPaymentPO(a, t, c string, b, f int, tk, u string) *PaymentPO {
	return &PaymentPO{
		Id:       0,
		Account:  a,
		PType:    t,
		Curr:     c,
		Balance:  b,
		Frozen:   f,
		Token:    tk,
		Remark:   "",
		CreateAt: 0,
		CreateBy: u,
	}
}

func (p *PaymentPO) Create() error {
	if p == nil {
		return fmt.Errorf("payment instance is nil")
	}
	r := global.DB.Create(p)
	return r.Error
}
