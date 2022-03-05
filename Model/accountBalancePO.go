package Model

import (
	"fmt"
	"htSample/global"
)

var errorAccountBalanceNull = fmt.Errorf("accountBalance instance is nil")

type AccountBalancePO struct {
	Id      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code    string `json:"code"`
	Curr    string `json:"curr"`
	Balance int64  `json:"balance"` //以0.01元为单位
}

func (a AccountBalancePO) TableName() string {
	return "account_balance"
}

func FindAccountBalanceAll() ([]AccountBalancePO, error) {
	r := make([]AccountBalancePO, 0)
	t := global.DB.Find(&r)
	return r, t.Error
}

func FindAccountBalancePOByCode(code string) (*AccountBalancePO, error) {
	r := &AccountBalancePO{}
	t := global.DB.Where("code = ?", code).First(r)
	return r, t.Error
}

func (a *AccountBalancePO) Save() error {
	if a == nil {
		return errorAccountBalanceNull
	}
	r := global.DB.Save(a)
	return r.Error
}
