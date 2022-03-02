package Model

type PaymentProcessPO struct {
	Id       int    `gorm:"id"`
	ACode    string `gorm:"acode"`
	TCode    string `gorm:"tcode"`   // sys_dict_type:1
	CCode    string `gorm:"ccode"`   // sys_dict_type:2
	Balance  int    `gorm:"balance"` // 以0.01元为单位
	Frozen   int    `gorm:"frozen"`  // 以0.01元为单位
	PToken   string `gorm:"ptoken"`
	Seq      int    `gorm:"seq"`
	Remark   string `gorm:"remark"`
	CreateAt int    `gorm:"create_at"`
	CreateBy string `gorm:"create_by"`
}

func (a PaymentProcessPO) TableName() string {
	return "payment_process"
}
