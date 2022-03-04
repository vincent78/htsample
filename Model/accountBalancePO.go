package Model

import "htSample/global"

type AccountBalancePO struct {
	Id      int    `gorm:"primaryKey;autoIncrement"`
	Code    string ``
	Curr    string ``
	Balance int64  `` //以0.01元为单位
	Frozen  int64  `` //以0.01元为单位
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
	t := global.DB.Where("code = ?", code).Find(r)
	return r, t.Error
}
