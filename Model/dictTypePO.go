package Model

import "htSample/global"

type DictTypePO struct {
	Id     int    `gorm:"primaryKey"`
	Code   string ``
	Name   string ``
	Status int    `gorm:"default:0"`
	Remark string ``
}

func (d DictTypePO) TableName() string {
	return "sys_dict_type"
}

func FindAllDictType() []DictTypePO {
	r := make([]DictTypePO, 0)
	global.DB.Where("status = 0").Find(&r)
	return r
}

func FindDictTypeByCode(code string) *DictTypePO {
	r := &DictTypePO{}
	global.DB.Where("status = 0 and code = ?", code).Find(r)
	return r
}
