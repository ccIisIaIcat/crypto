package query_bar_custom

import (
	"encoding/json"
	"global"
	"log"
	"strconv"
	"sync"
	"websocketlocal"
)

type QueryBar struct {
	// public
	InsId_list    []string
	Bar_info_chan chan *global.BarInfo // 对外提供的chan访问接口
	Custom_type   string
	// private
	local_ws         *websocketlocal.WebSocketLocal
	local_insid_info map[string]*global.BarInfo
	stop_signal      sync.Map
}

func (Q *QueryBar) Close() string {
	Q.stop_signal.Store("stop", true)
	Q.local_ws.Close()
	return "closed"
}

func (Q *QueryBar) init() {
	if len(Q.InsId_list) == 0 {
		panic("missing InsId")
	}
	if Q.Custom_type == "" {
		panic("missing custom type")
	}
	Q.local_insid_info = make(map[string]*global.BarInfo, 0)
	for i := 0; i < len(Q.InsId_list); i++ {
		Q.local_insid_info[Q.InsId_list[i]] = &global.BarInfo{}
		Q.local_insid_info[Q.InsId_list[i]].Insid = Q.InsId_list[i]
	}
	Q.stop_signal = sync.Map{}
	Q.stop_signal.Store("stop", false)
	Q.Bar_info_chan = make(chan *global.BarInfo, 1000)
	Q.local_ws = websocketlocal.GenWebSocket("wss://ws.okx.com:8443/ws/v5/public", 10)
}

// 对对应信息进行订阅(订阅bar信息)
func (Q *QueryBar) submit_bar() {
	for i := 0; i < len(Q.InsId_list); i++ {
		// 订阅bar信息
		Q.local_ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "candle`+Q.Custom_type+`","instId": "`+Q.InsId_list[i]+`"}]}`), true)
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
	// channel_info := temp["arg"].(map[string](interface{}))["channel"].(string)
	insid_info := temp["arg"].(map[string](interface{}))["instId"].(string)

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
	// signal, _ := temp["data"].([]interface{})[0].([]interface{})[8].(string)
	Q.local_insid_info[insid_info].Ts_open = ts_open
	Q.local_insid_info[insid_info].Open_price = open_price
	Q.local_insid_info[insid_info].High_price = high_price
	Q.local_insid_info[insid_info].Low_price = low_price
	Q.local_insid_info[insid_info].Close_price = close_price
	Q.local_insid_info[insid_info].Vol = vol
	Q.local_insid_info[insid_info].VolCcy = volCcy
	Q.local_insid_info[insid_info].VolCcyQuote = volCcyQuote
	// fmt.Println(Q.local_insid_info[insid_info])
	// fmt.Println(time.Now())
	Q.Bar_info_chan <- Q.local_insid_info[insid_info]

}

func (Q *QueryBar) Start() {
	Q.init()
	Q.submit_bar()
	go Q.local_ws.StartGather()
	judge, _ := Q.stop_signal.Load("stop")
	for !judge.(bool) {
		info := <-Q.local_ws.InfoChan
		// fmt.Println(string(info))
		Q.update_tick_info(info)
		judge, _ = Q.stop_signal.Load("stop")
	}
}

// func (Q *)
