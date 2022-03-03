package Model

type AccountBalancePO struct {
	Id      int    `gorm:"id;primaryKey"`
	Code    string ``
	Curr    string ``
	Balance int    `` //以0.01元为单位
	Frozen  int    `` //以0.01元为单位
}

func (a AccountBalancePO) TableName() string {
	return "account_balance"
}
