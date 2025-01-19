package api

import (
	"LotteryFilterAssistant/internal/entity"
)

// 参数结构体
type QueryOptions struct {
	Name     string // 名称
	PageNo   int    // 页码
	PageSize int    // 每页数量
}

type FucaiQuery interface {
	QueryFucai(options QueryOptions) (*entity.Fucai, error)
}
