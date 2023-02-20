package global

// 用于记录不同insid最新的信息
type TickInfo struct {
	// 量价信息
	Insid       string
	Ts_Price    int
	Ask1_price  float64
	Bid1_price  float64
	Ask1_volumn float64
	Bid1_volumn float64
	// open-interest信息
	Oi              int     // 持仓量，按张为单位
	OiCcy           float64 // 持仓量，按币为单位
	Ts_OpenInterest int     // open-interest最后更新时间
	// 资金费率信息
	FundingRate        float64 // 当前资金费率
	NextFundingRate    float64 // 下一期预测资金费率
	Ts_FundingRate     int     // 资金费率最后更新时间
	TS_NextFundingRate int
}
