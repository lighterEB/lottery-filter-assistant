package entity

import "strings"

// 响应结构体
type Response[T any] struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// 默认值
const (
	SUCCESS_CODE        = 200
	FAIL_CODE           = 500
	DEFAULT_MSG_SUCCESS = "请求成功"
	DEFAULT_MSG_FAILED  = "请求失败"
)

func ResSuccess[T any](data T) *Response[T] {
	return &Response[T]{
		Code:    SUCCESS_CODE,
		Success: true,
		Message: DEFAULT_MSG_SUCCESS,
		Data:    data,
	}
}

func ResFailed[T any](msg ...string) *Response[T] {
	message := DEFAULT_MSG_FAILED
	if len(msg) > 0 {
		message = strings.Join(msg, "")
	}
	return &Response[T]{
		Code:    FAIL_CODE,
		Success: false,
		Message: message,
	}
}
