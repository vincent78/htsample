package Model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAccountAll(t *testing.T) {
	err := ConnectDB(t)
	if err != nil {
		return
	}

	l := FindAccountAll()
	assert.Equal(t, len(l), 2)
}

func TestFindAccountByCode(t *testing.T) {
	err := ConnectDB(t)
	if err != nil {
		return
	}

	a := FindAccountByCode("a001")
	assert.Equal(t, a, &AccountPO{
		Id:      1,
		Code:    "a001",
		Name:    "bob123",
		Balance: 10000,
		Curr:    "usd",
		Remark:  "",
	})
}
