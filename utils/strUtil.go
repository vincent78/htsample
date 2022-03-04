package utils

import (
	"github.com/sony/sonyflake"
	"strconv"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{})

func GetUUID() string {
	id, _ := sf.NextID()
	return strconv.FormatUint(id, 16)
}
