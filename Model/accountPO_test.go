package Model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAccountAll(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	_, err = FindAccountAll()
	assert.NoError(t, err)
}

func TestFindAccountByCode(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	a, err := FindAccountByCode("a001")
	assert.NoError(t, err)
	assert.Equal(t, a, &AccountPO{
		Id:     1,
		Code:   "a001",
		Name:   "bob123",
		Remark: "",
	})
}
