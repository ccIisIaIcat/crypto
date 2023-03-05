package query_depth

import (
	"encoding/json"
	"global"
	"log"
	"strconv"
	"sync"
	"time"
	"websocketlocal"
)

type Query_depth struct {
	// public
	InsId_list     []string
	Tick_info_chan chan *global.Depth5Info // 对外提供的chan访问接口
	// private
	local_ws         *websocketlocal.WebSocketLocal
	local_insid_info map[string]*global.Depth5Info
	stop_signal      sync.Map
}

func (Q *Query_depth) init() {
	if len(Q.InsId_list) == 0 {
		panic("missing InsId")
	}
	Q.local_insid_info = make(map[string]*global.Depth5Info, 0)
	for i := 0; i < len(Q.InsId_list); i++ {
		Q.local_insid_info[Q.InsId_list[i]] = &global.Depth5Info{}
		Q.local_insid_info[Q.InsId_list[i]].Insid = Q.InsId_list[i]
	}
	Q.stop_signal = sync.Map{}
	Q.stop_signal.Store("stop", false)
	Q.Tick_info_chan = make(chan *global.Depth5Info, 1000)
	Q.local_ws = websocketlocal.GenWebSocket("wss://ws.okx.com:8443/ws/v5/public", 10)

}

// 对对应信息进行订阅(订阅bar信息)
func (Q *Query_depth) submit_tick() {
	for i := 0; i < len(Q.InsId_list); i++ {
		// 订阅bar信息
		Q.local_ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "books5","instId": "`+Q.InsId_list[i]+`"}]}`), true)
	}
}

func (Q *Query_depth) update_tick_info(temp_json []byte) {
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
		ts_price, _ := strconv.Atoi(temp["data"].([]interface{})[0].(map[string]interface{})["ts"].(string))
		temp_pointer := &global.Depth5Info{Insid: insid_info, Ts_Price: ts_price}
		asks1p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[0].([]interface{})[0].(string), 64)
		asks1q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[0].([]interface{})[1].(string), 64)
		asks2p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[1].([]interface{})[0].(string), 64)
		asks2q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[1].([]interface{})[1].(string), 64)
		asks3p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[2].([]interface{})[0].(string), 64)
		asks3q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[2].([]interface{})[1].(string), 64)
		asks4p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[3].([]interface{})[0].(string), 64)
		asks4q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[3].([]interface{})[1].(string), 64)
		asks5p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[4].([]interface{})[0].(string), 64)
		asks5q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["asks"].([]interface{})[4].([]interface{})[1].(string), 64)
		bids1p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[0].([]interface{})[0].(string), 64)
		bids1q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[0].([]interface{})[1].(string), 64)
		bids2p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[1].([]interface{})[0].(string), 64)
		bids2q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[1].([]interface{})[1].(string), 64)
		bids3p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[2].([]interface{})[0].(string), 64)
		bids3q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[2].([]interface{})[1].(string), 64)
		bids4p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[3].([]interface{})[0].(string), 64)
		bids4q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[3].([]interface{})[1].(string), 64)
		bids5p, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[4].([]interface{})[0].(string), 64)
		bids5q, _ := strconv.ParseFloat(temp["data"].([]interface{})[0].(map[string]interface{})["bids"].([]interface{})[4].([]interface{})[1].(string), 64)

		temp_pointer.Ask1_price = asks1p
		temp_pointer.Ask1_volumn = asks1q
		temp_pointer.Ask2_price = asks2p
		temp_pointer.Ask2_volumn = asks2q
		temp_pointer.Ask3_price = asks3p
		temp_pointer.Ask3_volumn = asks3q
		temp_pointer.Ask4_price = asks4p
		temp_pointer.Ask4_volumn = asks4q
		temp_pointer.Ask5_price = asks5p
		temp_pointer.Ask5_volumn = asks5q

		temp_pointer.Bid1_price = bids1p
		temp_pointer.Bid1_volumn = bids1q
		temp_pointer.Bid2_price = bids2p
		temp_pointer.Bid2_volumn = bids2q
		temp_pointer.Bid3_price = bids3p
		temp_pointer.Bid3_volumn = bids3q
		temp_pointer.Bid4_price = bids4p
		temp_pointer.Bid4_volumn = bids4q
		temp_pointer.Bid5_price = bids5p
		temp_pointer.Bid5_volumn = bids5q

		temp_pointer.Ts_Price = ts_price
		Q.Tick_info_chan <- temp_pointer

	}
}

func (Q *Query_depth) Start() {
	Q.init()
	Q.submit_tick()
	time.Sleep(time.Second)
	go Q.local_ws.StartGather()
	judge, _ := Q.stop_signal.Load("stop")
	for !judge.(bool) {
		info := <-Q.local_ws.InfoChan
		Q.update_tick_info(info)
		judge, _ = Q.stop_signal.Load("stop")
	}
}
