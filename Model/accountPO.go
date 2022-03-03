package Model

import "htSample/global"

type AccountPO struct {
	Id      int    `gorm:"primaryKey"`
	Code    string ``
	Name    string ``
	Balance int    ``
	Curr    string `` // sys_dict_type:2
	Remark  string ``
}

func (a AccountPO) TableName() string {
	return "account"
}

func FindAccountAll() []AccountPO {
	r := make([]AccountPO, 0)
	global.DB.Find(&r)
	return r
}

func FindAccountByCode(code string) *AccountPO {
	r := &AccountPO{}
	global.DB.Debug().Where("code = ?", code).Find(r)
	return r
}
