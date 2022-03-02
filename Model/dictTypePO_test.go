package Model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictTypeFindAll(t *testing.T) {
	err := ConnectDB(t)
	if err != nil {
		return
	}
	l := FindAllDictType()
	assert.Equal(t, len(l), 2)
}

func TestDictTypeFindCode(t *testing.T) {
	err := ConnectDB(t)
	if err != nil {
		return
	}

	o := FindDictTypeByCode("payment_type")
	assert.EqualValues(t, o, &DictTypePO{
		Id:     1,
		Code:   "payment_type",
		Name:   "交易类型",
		Status: 0,
		Remark: "",
	})
}
