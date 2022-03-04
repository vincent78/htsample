package Model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAccountBalanceAll(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	_, err = FindAccountBalanceAll()
	assert.NoError(t, err)
}

func TestFindAccountBalanceByCode(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	a, err := FindAccountBalancePOByCode("a001")
	assert.NoError(t, err)
	assert.Equal(t, a, &AccountBalancePO{
		Id:      1,
		Code:    "a001",
		Curr:    "usd",
		Balance: 10000,
		Frozen:  0,
	})
}
