package Model

import "htSample/global"

type AccountBalancePO struct {
	Id      int    `gorm:"primaryKey;autoIncrement"`
	Code    string ``
	Curr    string ``
	Balance int    `` //以0.01元为单位
	Frozen  int    `` //以0.01元为单位
}

func (a AccountBalancePO) TableName() string {
	return "account_balance"
}

func FindAccountBalanceAll() []AccountBalancePO {
	r := make([]AccountBalancePO, 0)
	global.DB.Find(&r)
	return r
}

func FindAccountBalancePOByCode(code string) *AccountBalancePO {
	r := &AccountBalancePO{}
	global.DB.Where("code = ?", code).Find(r)
	return r
}
