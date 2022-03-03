package Model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAccountBalanceAll(t *testing.T) {
	err := ConnectDB(t)
	if err != nil {
		return
	}

	l := FindAccountBalanceAll()
	assert.Equal(t, len(l), 2)
}

func TestFindAccountBalanceByCode(t *testing.T) {
	err := ConnectDB(t)
	if err != nil {
		return
	}

	a := FindAccountBalancePOByCode("a001")
	assert.Equal(t, a, &AccountBalancePO{
		Id:      1,
		Code:    "a001",
		Curr:    "usd",
		Balance: 10000,
		Frozen:  0,
	})
}
