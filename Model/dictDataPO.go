package Model

import "htSample/global"

type DictDataPO struct {
	Id     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Pid    int    `json:"pid"`
	Seq    int    `json:"seq"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Ext1   string `json:"ext1"`
	Ext2   string `json:"ext2"`
	Ext3   string `json:"ext3"`
	Status int    `json:"status"`
	Remark string `json:"remark"`
}

func (d DictDataPO) TableName() string {
	return "sys_dict_data"
}

func FindDictDataByPID(pid int) ([]DictDataPO, error) {
	r := make([]DictDataPO, 0)
	o := global.DB.Where("status = 0 and pid = ? ", pid).Order("seq asc").Find(&r)
	return r, o.Error
}

func FindDictDataListByPIDAndCode(pid int, code string) (*DictDataPO, error) {
	r := &DictDataPO{}
	o := global.DB.Where("status = 0 and pid = ? and code = ?", pid, code).Find(r)
	return r, o.Error
}
