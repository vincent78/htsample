package Model

type AccountBalancePO struct {
	Id      int    `gorm:"id"`
	ACode   string `gorm:"acode"`
	CCode   string `gorm:"ccode"`
	Balance int    `gorm:"balance"` //以0.01元为单位
	Frozen  int    `gorm:"frozen"`  //以0.01元为单位
}

func (a AccountBalancePO) TableName() string {
	return "account_balance"
}
