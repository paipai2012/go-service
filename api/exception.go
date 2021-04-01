package api

import (
	"encoding/json"
)

var _ Error = (*Exception)(nil)

type Error interface {
	// WithErrMsg 设置错误信息
	WithErrMsg(data interface{}) Error
	// ToString 返回 JSON 格式的错误详情
	ToString() string
}

type Exception struct {
	Code    int         `json:"code"`    // 业务编码
	Message string      `json:"message"` // 错误描述
	ErrMsg  interface{} `json:"errMsg"`  // 成功时返回的数据
}

func NewException(code int, message string) Error {
	return &Exception{
		Code:    code,
		Message: message,
		ErrMsg:  nil,
	}
}

func (e *Exception) WithErrMsg(errMsg interface{}) Error {
	e.ErrMsg = errMsg
	return e
}

func (e *Exception) ToString() string {
	json, _ := json.Marshal(e)
	return string(json)
}
