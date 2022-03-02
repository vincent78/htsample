package Model

import "htSample/global"

type DictDataPO struct {
	Id     int    `gorm:"id"`
	Pid    int    `gorm:"pid"`
	Seq    int    `gorm:"seq"`
	Code   string `gorm:"code"`
	Name   string `gorm:"name"`
	Ext1   string `gorm:"ext1"`
	Ext2   string `gorm:"ext2"`
	Ext3   string `gorm:"ext3"`
	Status int    `gorm:"status"`
	Remark string `gorm:"remark"`
}

func (d DictDataPO) TableName() string {
	return "sys_dict_data"
}

func FindDictDataByPID(pid int) []DictDataPO {
	r := make([]DictDataPO, 0)
	global.DB.Where("status = 0 and pid = ? ", pid).Order("seq asc").Find(&r)
	return r
}

func FindDictDataListByPIDAndCode(pid int, code string) *DictDataPO {
	r := &DictDataPO{}
	global.DB.Where("status = 0 and pid = ? and code = ?", pid, code).Find(r)
	return r
}
