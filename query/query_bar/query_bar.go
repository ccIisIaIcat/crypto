package query_bar

import (
	"encoding/json"
	"global"
	"log"
	"strconv"
	"websocketlocal"
)

type QueryBar struct {
	// public
	InsId_list    []string
	Bar_info_chan chan *global.BarInfo // 对外提供的chan访问接口
	InsType       string               // 暂时有三种，SPOT,SWAP,FUTURES
	// private
	local_ws         *websocketlocal.WebSocketLocal
	local_insid_info map[string]*global.BarInfo
}

func (Q *QueryBar) init() {
	if len(Q.InsId_list) == 0 {
		panic("missing InsId")
	}
	if Q.InsType == "" {
		panic("missing InsType")
	}
	Q.local_insid_info = make(map[string]*global.BarInfo, 0)
	for i := 0; i < len(Q.InsId_list); i++ {
		Q.local_insid_info[Q.InsId_list[i]] = &global.BarInfo{}
		Q.local_insid_info[Q.InsId_list[i]].Insid = Q.InsId_list[i]
	}
	Q.Bar_info_chan = make(chan *global.BarInfo, 1000)
	Q.local_ws = websocketlocal.GenWebSocket("wss://ws.okx.com:8443/ws/v5/public", 10)
}

// 对对应信息进行订阅(订阅bar信息)
func (Q *QueryBar) submit_bar() {
	for i := 0; i < len(Q.InsId_list); i++ {
		// 订阅bar信息
		Q.local_ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "candle1m","instId": "`+Q.InsId_list[i]+`"}]}`), true)
	}
}

// 对对应信息进行订阅(订阅资金费率信息)
func (Q *QueryBar) submit_fundingrate() {
	for i := 0; i < len(Q.InsId_list); i++ {
		// 订阅资金费率信息
		Q.local_ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "funding-rate","instId": "`+Q.InsId_list[i]+`"}]}`), true)
	}
}

// 对对应信息进行订阅(订阅持仓量信息)
func (Q *QueryBar) submit_openinterest() {
	for i := 0; i < len(Q.InsId_list); i++ {
		// 订阅资金费率信息
		Q.local_ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "open-interest","instId": "`+Q.InsId_list[i]+`"}]}`), true)
	}
}

func (Q *QueryBar) update_tick_info(temp_json []byte) {
	var temp map[string](interface{})
	err := json.Unmarshal(temp_json, &temp)
	if err != nil {
		log.Println("json解析错误", err)
		return
	}
	if temp["event"] != nil {
		return
	}
	channel_info := temp["arg"].(map[string](interface{}))["channel"].(string)
	insid_info := temp["arg"].(map[string](interface{}))["instId"].(string)

	if channel_info == "candle1m" {
		if temp["data"].([]interface{})[0].([]interface{})[8].(string) == "0" {
			return
		}
		ts_open, _ := strconv.Atoi(temp["data"].([]interface{})[0].([]interface{})[0].(string))
		open_price, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].([]interface{})[1].(string), 64)
		high_price, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].([]interface{})[2].(string), 64)
		low_price, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].([]interface{})[3].(string), 64)
		close_price, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].([]interface{})[4].(string), 64)
		vol, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].([]interface{})[5].(string), 64)
		volCcy, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].([]interface{})[6].(string), 64)
		volCcyQuote, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].([]interface{})[7].(string), 64)
		Q.local_insid_info[insid_info].Ts_open = ts_open
		Q.local_insid_info[insid_info].Open_price = open_price
		Q.local_insid_info[insid_info].High_price = high_price
		Q.local_insid_info[insid_info].Low_price = low_price
		Q.local_insid_info[insid_info].Close_price = close_price
		Q.local_insid_info[insid_info].Vol = vol
		Q.local_insid_info[insid_info].VolCcy = volCcy
		Q.local_insid_info[insid_info].VolCcyQuote = volCcyQuote
		Q.Bar_info_chan <- Q.local_insid_info[insid_info]
	} else if channel_info == "funding-rate" {
		fundingRate, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["fundingRate"].(string), 64)
		ts_fundingrate, _ := strconv.Atoi(temp["data"].([]interface{})[0].(map[string]interface{})["fundingTime"].(string))
		nextFundingRate, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["nextFundingRate"].(string), 64)
		nextFundingTime, _ := strconv.Atoi(temp["data"].([]interface{})[0].(map[string]interface{})["nextFundingTime"].(string))
		Q.local_insid_info[insid_info].FundingRate = fundingRate
		Q.local_insid_info[insid_info].Ts_FundingRate = ts_fundingrate
		Q.local_insid_info[insid_info].NextFundingRate = nextFundingRate
		Q.local_insid_info[insid_info].TS_NextFundingRate = nextFundingTime
	} else if channel_info == "open-interest" {
		oi, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["oi"].(string), 64)
		oiccy, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["oiCcy"].(string), 64)
		ts_oi, _ := strconv.Atoi(temp["data"].([]interface{})[0].(map[string]interface{})["ts"].(string))
		Q.local_insid_info[insid_info].Oi = oi
		Q.local_insid_info[insid_info].OiCcy = oiccy
		Q.local_insid_info[insid_info].Ts_oi = ts_oi
	}

}

func (Q *QueryBar) Start() {
	Q.init()
	if Q.InsType == "SWAP" {
		Q.submit_fundingrate()
	}
	Q.submit_bar()
	Q.submit_openinterest()
	go Q.local_ws.StartGather()
	for {
		info := <-Q.local_ws.InfoChan
		Q.update_tick_info(info)
	}
}

// func main() {
// 	qb := QueryBar{InsId_list: []string{"ETH-USDT-SWAP"}, InsType: "SWAP"}
// 	go qb.Start()
// 	time.Sleep(time.Second)
// 	for {
// 		temp := <-qb.Bar_info_chan
// 		fmt.Println(temp)
// 	}
// }
