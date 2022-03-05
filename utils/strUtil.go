package utils

import (
	"encoding/json"
	"fmt"
	"github.com/sony/sonyflake"
	"strconv"
	"strings"
	"unsafe"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{})

func GetUUID() string {
	id, _ := sf.NextID()
	return strconv.FormatUint(id, 16)
}

// ToJsonStr 将对象转为json字符串
func ToJsonStr(obj interface{}) string {

	if str, ok := obj.(string); ok {
		return str
	}

	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Errorf("strUtil.ToJsonStr error:%v", err.Error())
		return ""
	}
	s := fmt.Sprintf("%+v", Bytes2String(b))
	r := strings.Replace(s, `\u003c`, "<", -1)
	r = strings.Replace(r, `\u003e`, ">", -1)
	return r
}

// Bytes2String 直接转换底层指针，两者指向的相同的内存，改一个另外一个也会变。
// 效率是string([]byte{})的百倍以上，且转换量越大效率优势越明显。
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
