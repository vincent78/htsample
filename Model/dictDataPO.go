package Model

import "htSample/global"

type DictDataPO struct {
	Id     int    `gorm:"primaryKey"`
	Pid    int    ``
	Seq    int    ``
	Code   string ``
	Name   string ``
	Ext1   string ``
	Ext2   string ``
	Ext3   string ``
	Status int    ``
	Remark string ``
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
