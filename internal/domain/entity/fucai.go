package entity

// 福彩实体
type Fucai struct {
	State string `json:"state"` // 状态
	Message string `json:"message"` // 消息
	PageNo int `json:"pageNo"` // 页码
	PageNum int `json:"pageNum"` // 总页数
	PageSize int `json:"pageSize"` // 每页数量
	Result []Result `json:"result"` // 结果
}

type Result struct {
	Code string `json:"code"` // 期号
	Date string `json:"date"` // 日期
	Name string `json:"name"` // 名称
	Red string `json:"red"` // 红球/号码
	Blue string `json:"blue"` // 蓝球/号码
	Week string `json:"week"` // 星期
}
