package Model

type PaymentPO struct {
	Id       int    `gorm:"primaryKey"`
	Account  string ``
	PType    string `gorm:"column:ptype"` // sys_dict_type:1
	Curr     string ``                    // sys_dict_type:2
	Balance  int    ``                    // 以0.01元为单位
	Frozen   int    ``                    // 以0.01元为单位
	Token    string ``
	Remark   string `gorm:"remark"`
	CreateAt int    `gorm:"column:create_at"`
	CreateBy string `gorm:"column:create_by"`
}

func (a PaymentPO) TableName() string {
	return "payment"
}
