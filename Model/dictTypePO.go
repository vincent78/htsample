package Model

import "htSample/global"

type DictTypePO struct {
	Id     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status int    `gorm:"default:0" json:"status"`
	Remark string `json:"remark"`
}

func (d DictTypePO) TableName() string {
	return "sys_dict_type"
}

func FindAllDictType() ([]DictTypePO, error) {
	r := make([]DictTypePO, 0)
	o := global.DB.Where("status = 0").Find(&r)
	return r, o.Error
}

func FindDictTypeByCode(code string) (*DictTypePO, error) {
	r := &DictTypePO{}
	o := global.DB.Where("status = 0 and code = ?", code).Find(r)
	return r, o.Error
}
