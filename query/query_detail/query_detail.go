package query_detail

import (
	"encoding/json"
	"global"
	"log"
	"strconv"
	"websocketlocal"
)

type QueryDetail struct {
	// public
	InsId_list     []string
	Tick_info_chan chan *global.TickInfo // 对外提供的chan访问接口
	// private
	local_ws         *websocketlocal.WebSocketLocal
	local_insid_info map[string]*global.TickInfo
}

func (Q *QueryDetail) init() {
	if len(Q.InsId_list) == 0 {
		panic("missing InsId")
	}
	Q.local_insid_info = make(map[string]*global.TickInfo, 0)
	for i := 0; i < len(Q.InsId_list); i++ {
		Q.local_insid_info[Q.InsId_list[i]] = &global.TickInfo{}
		Q.local_insid_info[Q.InsId_list[i]].Insid = Q.InsId_list[i]
	}
	Q.Tick_info_chan = make(chan *global.TickInfo, 1000)
	Q.local_ws = websocketlocal.GenWebSocket("wss://ws.okx.com:8443/ws/v5/public", 10)

}

// 对对应信息进行订阅
func (Q *QueryDetail) submit() {
	for i := 0; i < len(Q.InsId_list); i++ {
		// 订阅量价信息
		Q.local_ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "price-limit","instId": "`+Q.InsId_list[i]+`"}]}`), true)
		// 订阅资金费率信息
		Q.local_ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "funding-rate","instId": "`+Q.InsId_list[i]+`"}]}`), true)
		// 订阅openiinterest信息
		Q.local_ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "open-interest","instId": "`+Q.InsId_list[i]+`"}]}`), true)

	}
}

// 对收到的json进行解析并更新对应tick_info
func (Q *QueryDetail) update_tick_info(temp_json []byte) {
	var temp map[string](interface{})
	err := json.Unmarshal(temp_json, &temp)
	if err != nil {
		log.Println("json解析错误", err)
		return
	}
	channel_info := temp["arg"].(map[string](interface{}))["channel"].(string)
	insid_info := temp["arg"].(map[string](interface{}))["instId"].(string)
	if temp["event"] != nil {
		return
	}
	if channel_info == "price-limit" {
		buy_limit, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["buyLmt"].(string), 64)
		sell_limit, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["sellLmt"].(string), 64)
		ts_price, _ := strconv.Atoi(temp["data"].([]interface{})[0].(map[string]interface{})["ts"].(string))
		Q.local_insid_info[insid_info].Bid1_price = buy_limit
		Q.local_insid_info[insid_info].Ask1_price = sell_limit
		Q.local_insid_info[insid_info].Ts_Price = ts_price
	} else if channel_info == "open-interest" {
		oi, _ := strconv.Atoi(temp["data"].([]interface{})[0].(map[string]interface{})["oi"].(string))
		oiCcy, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["oiCcy"].(string), 64)
		ts_open_interest, _ := strconv.Atoi(temp["data"].([]interface{})[0].(map[string]interface{})["ts"].(string))
		Q.local_insid_info[insid_info].Oi = oi
		Q.local_insid_info[insid_info].OiCcy = oiCcy
		Q.local_insid_info[insid_info].Ts_OpenInterest = ts_open_interest
	} else if channel_info == "funding-rate" {
		fundingRate, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["fundingRate"].(string), 64)
		ts_fundingrate, _ := strconv.Atoi(temp["data"].([]interface{})[0].(map[string]interface{})["fundingTime"].(string))
		nextFundingRate, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["nextFundingRate"].(string), 64)
		nextFundingTime, _ := strconv.Atoi(temp["data"].([]interface{})[0].(map[string]interface{})["nextFundingTime"].(string))
		Q.local_insid_info[insid_info].FundingRate = fundingRate
		Q.local_insid_info[insid_info].Ts_FundingRate = ts_fundingrate
		Q.local_insid_info[insid_info].NextFundingRate = nextFundingRate
		Q.local_insid_info[insid_info].TS_NextFundingRate = nextFundingTime
	}
	Q.Tick_info_chan <- Q.local_insid_info[insid_info]
}

func (Q *QueryDetail) Start() {
	Q.init()
	Q.submit()
	go Q.local_ws.StartGather()
	for {
		info := <-Q.local_ws.InfoChan
		Q.update_tick_info(info)
	}

}

// func main() {
// 	a := QueryDetail{InsId_list: []string{"BTC-USDT-SWAP"}}
// 	go a.Start()
// 	aa := 1
// 	for {
// 		aa += 1
// 		a := <-a.Tick_info_chan
// 		log.Println(a)
// 	}

// }
