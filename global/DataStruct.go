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

type Depth5Info struct {
	// 信息
	Insid       string
	Ts_Price    int
	Ask1_price  float64
	Ask2_price  float64
	Ask3_price  float64
	Ask4_price  float64
	Ask5_price  float64
	Bid1_price  float64
	Bid2_price  float64
	Bid3_price  float64
	Bid4_price  float64
	Bid5_price  float64
	Ask1_volumn float64
	Ask2_volumn float64
	Ask3_volumn float64
	Ask4_volumn float64
	Ask5_volumn float64
	Bid1_volumn float64
	Bid2_volumn float64
	Bid3_volumn float64
	Bid4_volumn float64
	Bid5_volumn float64
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
	// 该bar下最新持仓量信息
	Oi    float64
	OiCcy float64
	Ts_oi int
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

type SubmitInfo struct {
	Bar     *submitBar
	Tick    *submitTick
	Account *submitAccount
}

func GenSubmitInfo() *SubmitInfo {
	si := &SubmitInfo{}
	si.Bar = &submitBar{}
	si.Tick = &submitTick{}
	si.Account = &submitAccount{}
	return si
}

type submitBar struct {
	Judge       bool
	Port        string
	InsList     []string
	Custom_type string
}

type submitTick struct {
	Judge   bool
	Port    string
	InsList []string
}

type submitAccount struct {
	Judge         bool
	OrderJudge    bool
	PositionJudge bool
	AccountJudge  bool
	Simulate      bool
	Port          string
	Userconf      ConfigUser
}

// 当前所需信息：ctVal合约面值，ctMult合约乘数，ctValCcy合约面值计价货币，tickSz下单价格精度，lotSz下单数量精度，minSz最小下单数量，maxLmtSz合约或现货限价单的单笔最大委托数量,maxMktSz合约或现货市价单的单笔最大委托数量
type InsidBasicInfo struct {
	InsId    string `json:"Insid"`    // 合约名称
	CtVal    string `json:"ctVal"`    // 合约面值
	CtMult   string `json:"ctMult"`   // 合约乘数
	CtValCcy string `json:"ctValCcy"` // 合约面值计价货币，
	TickSz   string `json:"tickSz"`   // 下单价格精度
	LotSz    string `json:"lotSz"`    // 下单数量精度
	MinSz    string `json:"minSz"`    // 最小下单数量
	MaxLmtSz string `json:"maxLmtSz"` // 合约或现货限价单的单笔最大委托数量
	MaxMktSz string `json:"maxMktSz"` // 合约或现货市价单的单笔最大委托数量
}
