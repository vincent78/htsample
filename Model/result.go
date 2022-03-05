package Model

import (
	"fmt"
	"htSample/utils"
)
import "htSample/global"

type Result struct {
	Code    int         `json:"code"`  //当前状态
	Error   error       `json:"error"` //输出的内容(出错的提示）
	Context interface{} ``             //执行后的结果
}

type ResultOperate interface {
	Success() bool
	GetErrMsg() (int, string)
	GetContext()
}

func (r Result) Success() bool {
	return r.Code == global.SUCCESS
}

func (r Result) GetErrMsg() (int, string) {
	return r.Code, r.Error.Error()
}

func (r Result) GetContext() interface{} {
	return r.Context
}

func (r Result) ToString() string {
	return utils.ToJsonStr(r)
}

func (r Result) ToError() error {
	if r.Error != nil {
		return r.Error
	} else {
		return nil
	}
}

func SuccessResult(c interface{}) Result {
	return Result{
		Code:    global.SUCCESS,
		Context: c,
	}
}

func FailureResult(msg string, v ...interface{}) Result {
	return Result{
		Code:  global.FAILURE,
		Error: fmt.Errorf(msg, v),
	}
}
