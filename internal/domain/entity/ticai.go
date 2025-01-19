package entity

// 体彩实体
type Ticai struct {
	EmptyFlag    bool   `json:"emptyFlag"`    // 空标志
	ErrorCode    string `json:"errorCode"`    // 错误码
	ErrorMessage string `json:"errorMessage"` // 错误消息
	Success      bool   `json:"success"`      // 成功标志
	Value        Value  `json:"value"`        // 值
}

type Value struct {
	List     []ListItem `json:"list"`     // 列表
	PageNo   int        `json:"pageNo"`   // 页码
	PageSize int        `json:"pageSize"` // 每页数量
	Pages    int        `json:"pages"`    // 总页数
	Total    int        `json:"total"`    // 总数
}

type ListItem struct {
	DrawFlowFund      string `json:"drawFlowFund"`      // 开奖流水资金
	DrawPdfUrl        string `json:"drawPdfUrl"`        // 开奖PDF URL
	LotteryDrawNum    string `json:"lotteryDrawNum"`    // 期号
	LotteryDrawResult string `json:"lotteryDrawResult"` // 开奖结果
	LotteryDrawTime   string `json:"lotteryDrawTime"`   // 开奖时间
	LotteryGameName   string `json:"lotteryGameName"`   // 彩种名称
	LotteryGameNum    string `json:"lotteryGameNum"`    // 彩种编号
	LotteryUnsortDrawresult string `json:"lotteryUnsortDrawresult"` // 未排序开奖结果
}
