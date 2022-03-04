package Model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictTypeFindAll(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)
	_, err = FindAllDictType()
	assert.NoError(t, err)
}

func TestDictTypeFindCode(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	o, err := FindDictTypeByCode("payment_type")
	assert.NoError(t, err)
	assert.EqualValues(t, o, &DictTypePO{
		Id:     1,
		Code:   "payment_type",
		Name:   "交易类型",
		Status: 0,
		Remark: "",
	})
}
