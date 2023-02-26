package query_tick

import (
	"encoding/json"
	"global"
	"log"
	"strconv"
	"time"
	"websocketlocal"
)

type QueryTick struct {
	// public
	InsId_list     []string
	Tick_info_chan chan *global.TickInfo // 对外提供的chan访问接口
	InsType        string                // 暂时有三种，SPOT,SWAP,FUTURES
	// private
	local_ws         *websocketlocal.WebSocketLocal
	local_insid_info map[string]*global.TickInfo
}

func (Q *QueryTick) init() {
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

// 对对应信息进行订阅(订阅bar信息)
func (Q *QueryTick) submit_tick() {
	for i := 0; i < len(Q.InsId_list); i++ {
		// 订阅bar信息
		Q.local_ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "books5","instId": "`+Q.InsId_list[i]+`"}]}`), true)
	}
}

func (Q *QueryTick) update_tick_info(temp_json []byte) {
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

	if channel_info == "books5" {
		asks1p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[0].([]interface{})[0].(string), 64)
		asks1q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[0].([]interface{})[1].(string), 64)
		bids1p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[0].([]interface{})[0].(string), 64)
		bids1q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[0].([]interface{})[1].(string), 64)
		ts_price, _ := strconv.Atoi(temp["data"].([]interface{})[0].(map[string]interface{})["ts"].(string))
		Q.local_insid_info[insid_info].Ask1_price = asks1p
		Q.local_insid_info[insid_info].Ask1_volumn = asks1q
		Q.local_insid_info[insid_info].Bid1_price = bids1p
		Q.local_insid_info[insid_info].Bid1_volumn = bids1q
		Q.local_insid_info[insid_info].Ts_Price = ts_price
		Q.Tick_info_chan <- Q.local_insid_info[insid_info]
	}
}

func (Q *QueryTick) Start() {
	Q.init()
	Q.submit_tick()
	time.Sleep(time.Second)
	go Q.local_ws.StartGather()
	for {
		info := <-Q.local_ws.InfoChan
		Q.update_tick_info(info)
	}
}

// func main() {
// 	tick_chan := make(chan *global.TickInfo, 200)
// 	qt := QueryTick{InsId_list: []string{"BTC-USDT-SWAP"}, Tick_info_chan: tick_chan}
// 	go qt.Start()
// 	time.Sleep(time.Second)
// 	for {
// 		temp := <-qt.Tick_info_chan
// 		fmt.Println(temp)
// 	}
// }
