package Model

type AccountPO struct {
	Id      int    `gorm:"id"`
	Code    string `gorm:"code"`
	Name    string `gorm:"name"`
	Balance int    `gorm:"balance"`
	CCode   string `gorm:"ccode"` // sys_dict_type:2
	Remark  string `gorm:"remark"`
}

func (a AccountPO) TableName() string {
	return "account"
}
