package Model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDictDataByPID(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	_, err = FindDictDataByPID(1)
	assert.NoError(t, err)
}

func TestFindDictDataListByPIDAndCode(t *testing.T) {
	err := ConnectDB(t)
	assert.NoError(t, err)

	o, err := FindDictDataListByPIDAndCode(1, "outgoing")
	assert.NoError(t, err)
	assert.Equal(t, o, &DictDataPO{
		Id:     2,
		Pid:    1,
		Seq:    2,
		Code:   "outgoing",
		Name:   "outgoing",
		Ext1:   "",
		Ext2:   "",
		Ext3:   "",
		Status: 0,
		Remark: "出帐",
	})
}
