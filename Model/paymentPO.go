package Model

import (
	"fmt"
	"htSample/global"
)

var errorPaymentNull = fmt.Errorf("payment instance is nil")

type PaymentPO struct {
	Id       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Account  string `json:"account"`
	PType    string `gorm:"column:ptype" json:"ptype"` // sys_dict_type:1
	Curr     string `json:"curr"`                      // sys_dict_type:2
	Balance  int64  `json:"balance"`                   // 以0.01元为单位
	Frozen   int64  `json:"frozen"`                    // 以0.01元为单位
	Token    string `json:"token"`
	Remark   string `gorm:"remark" json:"remark"`
	CreateAt int    `gorm:"autoUpdateTime:milli" json:"createAt"`
	CreateBy string `gorm:"column:create_by" json:"createBy"`
}

func (p PaymentPO) TableName() string {
	return "payment"
}

func NewPaymentPO(a, t, c string, b, f int64, tk, u string) *PaymentPO {
	return &PaymentPO{
		Id:       0,
		Account:  a,
		PType:    t,
		Curr:     c,
		Balance:  b,
		Frozen:   f,
		Token:    tk,
		Remark:   "",
		CreateAt: 0,
		CreateBy: u,
	}
}
func FindPaymentPOList() ([]PaymentPO, error) {
	l := make([]PaymentPO, 0)
	r := global.DB.Order("create_at desc").Find(&l)
	return l, r.Error
}

func FindPaymentPOListByAccount(a string) ([]PaymentPO, error) {
	l := make([]PaymentPO, 0)
	r := global.DB.Where("account = ?", a).Order("create_at desc").Find(&l)
	return l, r.Error
}

func FindPaymentPOListByToken(t string) ([]PaymentPO, error) {
	r := make([]PaymentPO, 0)
	o := global.DB.Where("token = ?", t).Order("ptype desc").Find(&r)
	return r, o.Error
}

func FindPaymentPOListByAccountCurr(a, c string) (*PaymentPO, error) {
	r := &PaymentPO{}
	o := global.DB.Where("account = ? and curr = ?", a, c).Find(&r)
	return r, o.Error
}

func FindPaymentPOByAccountToken(a, t string) (*PaymentPO, error) {
	r := &PaymentPO{}
	o := global.DB.Where("account = ? and token = ?", a, t).Find(r)
	return r, o.Error
}

func (p *PaymentPO) Create() error {
	if p == nil {
		return errorPaymentNull
	}
	r := global.DB.Create(p)
	return r.Error
}

func (p *PaymentPO) Save() error {
	if p == nil {
		return errorPaymentNull
	}
	r := global.DB.Save(p)
	return r.Error
}

func (p *PaymentPO) refreshNum() error {
	if p == nil {
		return errorPaymentNull
	}
	t, e := FindPaymentPOByAccountToken(p.Account, p.Token)
	if e != nil {
		return e
	}
	if t == nil {
		return fmt.Errorf("not find the payment by token:%v", p.Token)
	}
	p.Balance = t.Balance
	p.Frozen = t.Frozen
	return nil
}

func (p *PaymentPO) FrozenNum(num int64) error {
	if e := p.refreshNum(); e != nil {
		return e
	}

	if p.Balance < num {
		return fmt.Errorf("account:%v balance(%v) < num(%v)", p.Account, p.Balance, num)
	}
	p.Balance -= num
	p.Frozen += num
	r := global.DB.Model(p).Updates(map[string]interface{}{"frozen": p.Frozen, "balance": p.Balance})
	return r.Error
}

func (p *PaymentPO) CommitNum(num int64) error {
	if e := p.refreshNum(); e != nil {
		return e
	}
	if p.Frozen < num {
		return fmt.Errorf("account:%v Frozen(%v) < num(%v)", p.Account, p.Frozen, num)
	}
	p.Frozen -= num
	p.Balance += num
	//r := global.DB.Model(p).Updates(map[string]interface{}{"frozen": p.Frozen})
	r := global.DB.Save(p)
	return r.Error
}
