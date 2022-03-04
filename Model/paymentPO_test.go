package Model

import (
	"github.com/stretchr/testify/assert"
	"htSample/utils"
	"testing"
)

func TestPaymentPO_Create(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	tk := utils.GetUUID()

	p := NewPaymentPO(
		"a001",
		"incoming",
		"usd",
		100,
		101,
		tk,
		"admin",
	)
	err = p.Create()
	assert.NoError(t, err, "no error")

	o, err := FindPaymentPOByToken(p.Account, p.Token)
	assert.NoError(t, err)
	assert.Equal(t, o.Frozen, int64(101))
}

func TestFindPaymentPOListByAccount(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	_, err = FindPaymentPOListByAccount("a001")
	assert.NoError(t, err)
}

func TestPaymentPO(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	tk := utils.GetUUID()

	const sb = int64(110)
	const sf = int64(0)
	const num = int64(31)
	pi := NewPaymentPO(
		"t001",
		"incoming",
		"usd",
		sb,
		sf,
		tk,
		"admin",
	)

	err = pi.Create()
	assert.NoError(t, err, "create error")

	po := NewPaymentPO(
		"t002",
		"outgoing",
		"usd",
		sb,
		sf,
		tk,
		"admin",
	)

	err = po.Create()
	assert.NoError(t, err, "create error")

	err = po.FrozenNum(num)
	assert.NoError(t, err, "frozen error")
	assert.Equal(t, po.Balance, sb-num, "the balance error")
	assert.Equal(t, po.Frozen, sf+num, "the frozen error")

	err = po.CommitNum(num)
	assert.NoError(t, err, "commit error")
	assert.Equal(t, po.Frozen, int64(0), "the frozen error")

}
