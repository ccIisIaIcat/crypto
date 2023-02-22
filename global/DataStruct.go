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

type BarInfo struct {
	// bar信息
	Insid       string
	Ts_open     int
	Open_price  float64
	High_price  float64
	Low_price   float64
	Close_price float64
	Vol         float64 // 交易量，以张为单位
	VolCcy      float64 // 交易量，以币为单位
	VolCcyQuote float64 // 交易量，以计价货币为单位
	// 如果是SWAP，还保存资金费率信息
	FundingRate        float64 // 当前资金费率
	NextFundingRate    float64 // 下一期预测资金费率
	Ts_FundingRate     int     // 资金费率最后更新时间
	TS_NextFundingRate int
}

type Config struct {
	MysqlInfo map[string]ConfigMysql
	UserInfo  map[string]ConfigUser
}

type ConfigMysql struct {
	Host     string
	Port     string
	User     string
	Password string
}

type ConfigUser struct {
	Apikey     string
	Secretkey  string
	Passphrase string
}
