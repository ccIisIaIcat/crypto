package trade_restful

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"global"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type TradeRestful struct {
	client   *http.Client
	userconf global.ConfigUser
	baseurl  string
}

func GenTradeRestful(userconf global.ConfigUser) *TradeRestful {
	var judge global.ConfigUser
	if judge == userconf {
		panic("missing userconf")
	}
	tr := TradeRestful{}
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

func (T *TradeRestful) GenSign(method string, requestPath string, body string) (string, string) {
	temp_ts := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	return temp_ts, ComputeHmacSha256(temp_ts+method+requestPath+body, T.userconf.Secretkey)
}

func (T *TradeRestful) SendOrder(body string) string {
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

func (T *TradeRestful) ChangeLeverage(InsId string, leverage string, mgnMode string) string {
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

// func main() {
// 	sample_json := `{"instId":"ETH-USDT-SWAP","posSide":"long","tdMode":"cross","side":"sell","ordType":"market","sz":"1","clOrdId":"sA01"}`
// 	conf := global.GetConfig("../../conf/conf.ini")
// 	tr := GenTradeRestful(conf.UserInfo["1"])
// 	lalala := tr.ChangeLeverage("ETH-USDT-SWAP", "5", "cross")
// 	log.Println(lalala)
// 	lala := tr.SendOrder(sample_json)
// 	log.Println(lala)
// 	var temp map[string]interface{}
// 	json.Unmarshal([]byte(lala), &temp)
// 	fmt.Println(temp["code"].(string))
// 	order_name := temp["data"].([]interface{})[0].(map[string]interface{})["clOrdId"]
// 	fmt.Println(strings.Split(order_name.(string), "0"))
// }

// func main() {
// 	user_conf := global.GetConfig("../../conf/conf.ini")
// 	gtrs := GenTradeRestfulSimulate(user_conf.UserInfo["Simulate"])
// 	order := `{"clOrdId":"sA32","instId": "BTC-USDT-SWAP","tdMode": "cross","posSide":"short","side": "sell","ordType": "market","sz": "100"}`
// 	// order := `{"instId":"ETH-USDT-SWAP","posSide":"","tdMode":"cross","side":"buy","ordType":"market","sz":"1"}`
// 	res := gtrs.SendOrder(order)
// 	fmt.Println(res)
// 	var temp map[string]interface{}
// 	json.Unmarshal([]byte(res), &temp)
// 	fmt.Println(temp["code"].(string))
// 	fmt.Println(temp["data"].([]interface{})[0].(map[string]interface{})["clOrdId"])
// }
