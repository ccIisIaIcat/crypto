package trade_restful_simulate

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

type TradeRestfulSimulate struct {
	client   *http.Client
	userconf global.ConfigUser
	baseurl  string
}

func GenTradeRestfulSimulate(userconf global.ConfigUser) *TradeRestfulSimulate {
	var judge global.ConfigUser
	if judge == userconf {
		panic("missing userconf")
	}
	tr := TradeRestfulSimulate{}
	tr.userconf = userconf
	tr.client = &http.Client{}
	tr.baseurl = "https://www.okx.com"
	return &tr
}

func ComputeHmacSha256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (T *TradeRestfulSimulate) GenSign(method string, requestPath string, body string) (string, string) {
	temp_ts := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	return temp_ts, ComputeHmacSha256(temp_ts+method+requestPath+body, T.userconf.Secretkey)
}

func (T *TradeRestfulSimulate) SendOrder(body string) string {
	body = simulateProcess(body)
	reqest, err := http.NewRequest("POST", T.baseurl+"/api/v5/trade/order", strings.NewReader(body))
	if err != nil {
		log.Println("请求错误")
		panic(err)
	}
	ts, sign := T.GenSign("POST", "/api/v5/trade/order", body)
	// OK-ACCESS-KEY
	reqest.Header.Add("OK-ACCESS-KEY", T.userconf.Apikey)
	reqest.Header.Add("OK-ACCESS-PASSPHRASE", T.userconf.Passphrase)
	reqest.Header.Add("OK-ACCESS-SIGN", sign)
	reqest.Header.Add("OK-ACCESS-TIMESTAMP", ts)
	reqest.Header.Add("content-type", "application/json")
	reqest.Header.Add("x-simulated-trading", "1")

	resp, err := T.client.Do(reqest)
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

func (T *TradeRestfulSimulate) ChangeLeverage(InsId string, leverage string, mgnMode string) string {
	temp_body := `{"instId":"` + InsId + `","lever":"` + leverage + `","mgnMode":"` + mgnMode + `"}`
	reqest, err := http.NewRequest("POST", T.baseurl+"/api/v5/account/set-leverage", strings.NewReader(temp_body))
	if err != nil {
		log.Println("请求错误")
		panic(err)
	}
	ts, sign := T.GenSign("POST", "/api/v5/account/set-leverage", temp_body)
	// OK-ACCESS-KEY
	reqest.Header.Add("OK-ACCESS-KEY", T.userconf.Apikey)
	reqest.Header.Add("OK-ACCESS-PASSPHRASE", T.userconf.Passphrase)
	reqest.Header.Add("OK-ACCESS-SIGN", sign)
	reqest.Header.Add("OK-ACCESS-TIMESTAMP", ts)
	reqest.Header.Add("content-type", "application/json")
	reqest.Header.Add("x-simulated-trading", "1")

	resp, err := T.client.Do(reqest)
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

// 模拟交易时的swap订单没有posSide选项，使对饮字符段为空并根据long、short确定side字段
func simulateProcess(original_order string) string {
	var temp map[string]interface{}
	json.Unmarshal([]byte(original_order), &temp)
	if temp["instId"].(string)[len(temp["instId"].(string))-4:] == "SWAP" {
		if temp["posSide"].(string) == "long" {
			temp["posSide"] = ""
			temp["side"] = "buy"
		} else {
			temp["posSide"] = ""
			temp["side"] = "sell"
		}
	} else {
		return original_order
	}
	an, _ := json.Marshal(temp)
	return string(an)
}

// func main() {
// 	user_conf := global.GetConfig("../../conf/conf.ini")
// 	gtrs := GenTradeRestfulSimulate(user_conf.UserInfo["Simulate"])
// 	order := `{"instId": "BTC-USDT-SWAP","tdMode": "cross","posSide":"short","side": "sell","ordType": "market","sz": "100"}`
// 	// order := `{"instId":"ETH-USDT-SWAP","posSide":"","tdMode":"cross","side":"buy","ordType":"market","sz":"1"}`
// 	res := gtrs.SendOrder(order)
// 	fmt.Println(res)
// }
