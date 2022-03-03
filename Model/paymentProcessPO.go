package Model

type PaymentProcessPO struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Account  string ``
	PType    string `gorm:"column:ptype"` // sys_dict_type:1
	Curr     string ``                    // sys_dict_type:2
	Balance  int    `gorm:"balance"`      // 以0.01元为单位
	Frozen   int    `gorm:"frozen"`       // 以0.01元为单位
	Token    string ``
	Seq      int    `gorm:"seq"`
	Remark   string `gorm:"remark"`
	CreateAt int    `gorm:"autoUpdateTime:milli"`
	CreateBy string `gorm:"column:create_by"`
}

func (a PaymentProcessPO) TableName() string {
	return "payment_process"
}
