package account

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"global"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// 该结构体完成策略初始化时的持仓模式设置，杠杆倍数设置，相关交易币种产品频道的基本信息，以及产品频道基本信息更改的行情推送
type AccountConf struct {
	userconf          global.ConfigUser
	restful_baseurl   string
	websocket_baseurl string
	client            *http.Client
	simulate          bool
}

func GenAccountConf(userconf global.ConfigUser, simulate bool) *AccountConf {
	acc := &AccountConf{}
	acc.userconf = userconf
	acc.client = &http.Client{}
	acc.simulate = simulate
	if simulate {
		acc.restful_baseurl = "https://aws.okx.com"
		acc.websocket_baseurl = "wss://wspap.okx.com:8443/ws/v5/public?brokerId=9999"
	} else {
		acc.restful_baseurl = "https://aws.okx.com"
		acc.websocket_baseurl = "wss: //ws.okx.com:8443/ws/v5/public"
	}

	return acc
}

// 设置某个品种的杠杆倍数
func (A *AccountConf) SetLeverage(InsId string, leverage string, mgnMode string) string {
	temp_body := `{"instId":"` + InsId + `","lever":"` + leverage + `","mgnMode":"` + mgnMode + `"}`
	reqest, err := http.NewRequest("POST", A.restful_baseurl+"/api/v5/account/set-leverage", strings.NewReader(temp_body))
	if err != nil {
		log.Println("请求错误")
		panic(err)
	}
	ts, sign := A.GenSign("POST", "/api/v5/account/set-leverage", temp_body)
	// OK-ACCESS-KEY
	reqest.Header.Add("OK-ACCESS-KEY", A.userconf.Apikey)
	reqest.Header.Add("OK-ACCESS-PASSPHRASE", A.userconf.Passphrase)
	reqest.Header.Add("OK-ACCESS-SIGN", sign)
	reqest.Header.Add("OK-ACCESS-TIMESTAMP", ts)
	reqest.Header.Add("content-type", "application/json")
	if A.simulate {
		reqest.Header.Add("x-simulated-trading", "1")
	}
	resp, err := A.client.Do(reqest)
	if err != nil {
		log.Println("发送错误")
		panic(err)
	}
	defer resp.Body.Close()
	body2, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("解析错误")
		panic(err)
	}
	return string(body2)
}

// 设置持仓模式
func (A *AccountConf) SetPositionMode(mode string) string {
	temp_body := ""
	if mode == "longshort" || mode == "long_short" || mode == "long_short_mode" || mode == "longshort_mode" {
		temp_body = `{"posMode":"long_short_mode"}`
	} else if mode == "net" || mode == "net_mode" {
		temp_body = `{"posMode":"net_mode"}`
	} else {
		return "set mode err, check spell or api info"
	}
	reqest, err := http.NewRequest("POST", A.restful_baseurl+"/api/v5/account/set-position-mode", strings.NewReader(temp_body))
	if err != nil {
		log.Println("请求错误")
		panic(err)
	}
	ts, sign := A.GenSign("POST", "/api/v5/account/set-leverage", temp_body)
	// OK-ACCESS-KEY
	reqest.Header.Add("OK-ACCESS-KEY", A.userconf.Apikey)
	reqest.Header.Add("OK-ACCESS-PASSPHRASE", A.userconf.Passphrase)
	reqest.Header.Add("OK-ACCESS-SIGN", sign)
	reqest.Header.Add("OK-ACCESS-TIMESTAMP", ts)
	reqest.Header.Add("content-type", "application/json")
	if A.simulate {
		reqest.Header.Add("x-simulated-trading", "1")
	}
	resp, err := A.client.Do(reqest)
	if err != nil {
		log.Println("发送错误")
		panic(err)
	}
	defer resp.Body.Close()
	body2, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("解析错误")
		panic(err)
	}
	return string(body2)
}

// 获取产品相关基础信息
// 当前所需信息：ctVal合约面值，ctMult合约乘数，ctValCcy合约面值计价货币，tickSz下单价格精度，lotSz下单数量精度，minSz最小下单数量，maxLmtSz合约或现货限价单的单笔最大委托数量,maxMktSz合约或现货市价单的单笔最大委托数量
func (A *AccountConf) GetInsIdInfo(InsType string, InsList []string) []byte {
	res, err := http.Get(A.restful_baseurl + "/api/v5/public/instruments?instType=" + InsType)
	if err != nil {
		log.Println(err)
	}
	tool_map := make(map[string]bool)
	for i := 0; i < len(InsList); i++ {
		tool_map[InsList[i]] = true
	}
	ans_map := make(map[string]*global.InsidBasicInfo)
	temp_json, _ := io.ReadAll(res.Body)
	var temp map[string]interface{}
	json.Unmarshal(temp_json, &temp)
	info_list := temp["data"].([]interface{})
	for i := 0; i < len(info_list); i++ {
		if _, ok := tool_map[info_list[i].(map[string]interface{})["instId"].(string)]; ok {
			temp_info := info_list[i].(map[string]interface{})
			temp_basic := &global.InsidBasicInfo{}
			temp_basic.InsId = temp_info["instId"].(string)
			temp_basic.CtVal = temp_info["ctVal"].(string)
			temp_basic.CtValCcy = temp_info["ctValCcy"].(string)
			temp_basic.TickSz = temp_info["tickSz"].(string)
			temp_basic.LotSz = temp_info["lotSz"].(string)
			temp_basic.MinSz = temp_info["minSz"].(string)
			temp_basic.MaxLmtSz = temp_info["maxLmtSz"].(string)
			temp_basic.MaxMktSz = temp_info["maxMktSz"].(string)
			ans_map[temp_basic.InsId] = temp_basic
		}
	}
	json_info, _ := json.Marshal(ans_map)
	return json_info
}

func ComputeHmacSha256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (A *AccountConf) GenSign(method string, requestPath string, body string) (string, string) {
	temp_ts := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	return temp_ts, ComputeHmacSha256(temp_ts+method+requestPath+body, A.userconf.Secretkey)
}
