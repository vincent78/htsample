package Model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDictDataByPID(t *testing.T) {
	err := ConnectDB(t)
	if err != nil {
		return
	}

	l := FindDictDataByPID(1)
	assert.Equal(t, len(l), 3)
}

func TestFindDictDataListByPIDAndCode(t *testing.T) {
	err := ConnectDB(t)
	if err != nil {
		return
	}

	o := FindDictDataListByPIDAndCode(1, "outgoing")
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
