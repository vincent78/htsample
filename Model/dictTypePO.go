package Model

import "htSample/global"

type DictTypePO struct {
	Id     int    `gorm:"id"`
	Code   string `gorm:"code"`
	Name   string `gorm:"name"`
	Status int    `gorm:"status"`
	Remark string `gorm:"remark"`
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
