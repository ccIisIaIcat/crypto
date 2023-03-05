package account

import (
	"encoding/json"
	"global"
	"log"
	"strconv"
	"sync"
	"time"
	"websocketlocal"
)

type Account struct {
	// public
	InfoChanAccount   chan []byte
	InfoChanOrders    chan []byte
	InfoChanPositions chan []byte
	// private
	ws             *websocketlocal.WebSocketLocal
	userconf       global.ConfigUser
	accountsignal  bool
	orderssiganl   bool
	positiossignal bool
	stop_signal    sync.Map
}

// simulate_api指定api是否为模拟api
func GenAccount(userconf global.ConfigUser, account_sub bool, order_sub bool, position_sub bool, simulate_api bool) *Account {
	ac := &Account{}
	if account_sub {
		ac.InfoChanAccount = make(chan []byte, 1000)
	}
	if order_sub {
		ac.InfoChanOrders = make(chan []byte, 1000)
	}
	if position_sub {
		ac.InfoChanPositions = make(chan []byte, 1000)
	}
	ac.accountsignal = account_sub
	ac.orderssiganl = order_sub
	ac.positiossignal = position_sub
	if simulate_api {
		ac.ws = websocketlocal.GenWebSocket("wss://wspap.okx.com:8443/ws/v5/private?brokerId=9999", 10)
	} else {
		ac.ws = websocketlocal.GenWebSocket("wss://ws.okx.com:8443/ws/v5/private", 10)
	}
	ac.userconf = userconf
	ac.stop_signal = sync.Map{}
	ac.stop_signal.Store("stop", false)
	ac.login()
	time.Sleep(time.Second)
	if account_sub {
		ac.subcribeAccount()
	}
	if order_sub {
		ac.subcribeOrder()
	}
	if position_sub {
		ac.subcribePosition()
	}

	return ac
}

func (A *Account) Close() {
	A.stop_signal.Store("stop", true)
	A.ws.Close()
}

func (A *Account) genSign() (string, string) {
	method := "GET"
	requestPath := "/users/self/verify"
	temp_ts := strconv.Itoa(int(time.Now().UTC().Unix()))
	return temp_ts, global.ComputeHmacSha256(temp_ts+method+requestPath, A.userconf.Secretkey)
}

func (A *Account) login() {
	tempts, sigh := A.genSign()
	login_info := `{"op": "login","args":[{"apiKey":"` + A.userconf.Apikey + `","passphrase" :"` + A.userconf.Passphrase + `","timestamp" :"` + tempts + `","sign" :"` + sigh + `" }]}`
	A.ws.Submit([]byte(login_info), true)
}

// 账户频道
// 获取账户信息，首次订阅按照订阅维度推送数据，此外，当下单、撤单、成交等事件触发时，推送数据以及按照订阅维度定时推送数据
func (A *Account) subcribeAccount() {
	A.ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "account"}]}`), true)
}

// 持仓频道
// 获取持仓信息，首次订阅按照订阅维度推送数据，此外，当下单、撤单等事件触发时，推送数据以及按照订阅维度定时推送数据
func (A *Account) subcribePosition() {
	A.ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "positions","instType": "ANY"}]}`), true)
}

// 订单频道
// 获取订单信息，首次订阅不推送，只有当下单、撤单等事件触发时，推送数据
func (A *Account) subcribeOrder() {
	A.ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "orders","instType": "ANY"}]}`), true)
}

func (A *Account) Start() {
	go A.ws.StartGather()
	time.Sleep(time.Second)
	judge, _ := A.stop_signal.Load("stop")
	for !judge.(bool) {
		temp := <-A.ws.InfoChan
		var tempmap map[string]interface{}
		json.Unmarshal(temp, &tempmap)
		if tempmap["event"] != nil {
			log.Println(string(temp))
		} else {
			if tempmap["arg"].(map[string]interface{})["channel"].(string) == "account" {
				A.InfoChanAccount <- temp
			} else if tempmap["arg"].(map[string]interface{})["channel"].(string) == "orders" {
				A.InfoChanOrders <- temp
			} else if tempmap["arg"].(map[string]interface{})["channel"].(string) == "positions" {
				A.InfoChanPositions <- temp
			}
		}
		judge, _ = A.stop_signal.Load("stop")
	}
}

// func main() {
// 	conf := global.GetConfig("../conf/conf.ini")
// 	ac := GenAccount(conf.UserInfo["1"], false, true, true)
// 	go ac.Start()
// 	time.Sleep(time.Second)
// 	for {
// 		temp := <-ac.InfoChanPositions
// 		log.Println(string(temp))
// 	}
// }
