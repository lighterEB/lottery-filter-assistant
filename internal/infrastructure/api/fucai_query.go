package api

import "LotteryFilterAssistant/internal/domain/entity"

// 参数结构体
type FucaiParams struct {
	Name     string // 名称
	PageNo   int    // 页码
	PageSize int    // 每页数量
}

const (
	BASE_URL = "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?"
)

type fucaiQuery interface {
	queryFucai(params FucaiParams) (*entity.Fucai, error)
}
