package Model

import "htSample/global"

type AccountPO struct {
	Id      int    `gorm:"primaryKey;autoIncrement"`
	Code    string ``
	Name    string ``
	Balance int    ``
	Curr    string `` // sys_dict_type:2
	Remark  string ``
}

func (a AccountPO) TableName() string {
	return "account"
}

func FindAccountAll() ([]AccountPO, error) {
	r := make([]AccountPO, 0)
	o := global.DB.Find(&r)
	return r, o.Error
}

func FindAccountByCode(code string) (*AccountPO, error) {
	r := &AccountPO{}
	o := global.DB.Where("code = ?", code).Find(r)
	return r, o.Error
}
