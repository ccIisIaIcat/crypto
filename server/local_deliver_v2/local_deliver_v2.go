package main

import (
	"encoding/json"
	"fmt"
	"global"
	deliver "godeliver"
	"local_deliver_v2/deliver_server"
	"strings"
	"time"
)

type InfoDeliver struct {
	submitserver_port    string
	orderserver_port     string
	userconfig           global.ConfigUser
	simulate             bool
	strategy_map         map[string]*deliver_server.StrategyUnit
	pingpong_map         map[string](chan bool)
	orderserver          *deliver_server.OrderServer
	account_info_deliver *deliver_server.Account_deliver
	time_out             int
}

// 赋值端口和参数
func GenInfoDeliver(submitserver_port string, orderserver_port string, userconfig global.ConfigUser, simulate_account bool, time_out int) *InfoDeliver {
	ifd := &InfoDeliver{}
	ifd.time_out = time_out
	ifd.strategy_map = make(map[string]*deliver_server.StrategyUnit, 0)
	ifd.pingpong_map = make(map[string]chan bool, 0)
	ifd.submitserver_port = submitserver_port
	ifd.orderserver_port = orderserver_port
	ifd.userconfig = userconfig
	ifd.simulate = simulate_account
	ifd.account_info_deliver = deliver_server.GenAccountDeliver(ifd.userconfig, ifd.simulate, ifd.time_out)
	return ifd
}

// 初始化服务，直接开启Orderserver,实时接收InfoDeliver收到的请求，进行对应操作
func (I *InfoDeliver) Start() {
	I.orderserver = &deliver_server.OrderServer{}
	go I.orderserver.OrderServerListen(I.orderserver_port, I.userconfig, I.simulate)
	submitserver := deliver_server.SubmitServer{Userconf: I.userconfig, Simulate: I.simulate}
	go submitserver.SubmitServerListen(I.submitserver_port)
	go I.OrderResDeliver()

	go I.account_info_deliver.DeliverAccount()
	fmt.Println("server starting")
	for {
		select {
		case temp := <-submitserver.InfoChan:
			if temp.Subtype != "Ping" && temp.Subtype != "ping" {
				fmt.Println("-------------new request-----------------")
				fmt.Println(temp)
				fmt.Println("-----------------------------------------")
				strategy_name, sub_info := I.DealRequest(temp)
				I.pingpong_map[strategy_name] = make(chan bool, 10)
				I.strategy_map[strategy_name] = deliver_server.GenStrategyUnit(strategy_name, I.time_out, *sub_info, I.pingpong_map[strategy_name], I.simulate)
				go func() {
					I.strategy_map[strategy_name].Start()
					delete(I.strategy_map, strategy_name)
					delete(I.pingpong_map, strategy_name)
				}()
				go func() {
					if sub_info.Account.Judge {
						I.account_info_deliver.AddStrategy(strategy_name, sub_info.Account.Port, sub_info.Account.AccountJudge, sub_info.Account.PositionJudge, sub_info.Account.OrderJudge)
						I.account_info_deliver.CancelStrategy(strategy_name)
					}
				}()
			} else {
				if _, ok := I.pingpong_map[temp.Strategyname]; ok {
					I.pingpong_map[temp.Strategyname] <- true
				}
				if _, ok := I.account_info_deliver.PingPongMapChan[temp.Strategyname]; ok {
					I.account_info_deliver.PingPongMapChan[temp.Strategyname] <- true
				}
			}
		case <-time.After(time.Millisecond * 200):
		}
	}
}

func (I *InfoDeliver) OrderResDeliver() {
	for {
		select {
		case res := <-I.orderserver.Res_chan:
			var temp map[string]interface{}
			json.Unmarshal([]byte(res), &temp)
			fmt.Println("////")
			fmt.Println(temp)
			if temp["code"].(string) != "0" {
				order_name := temp["data"].([]interface{})[0].(map[string]interface{})["clOrdId"].(string)
				strategy_name := strings.Split(order_name, "0")[0]
				if I.strategy_map[strategy_name].Submitinfo.Account.Judge && I.strategy_map[strategy_name].Submitinfo.Account.OrderJudge {
					error_info := `{"arg": {"channel": "orders",},"data": [{"clOrdId" : "` + order_name + `","state": "placed error","sMsg":"` + temp["data"].([]interface{})[0].(map[string]interface{})["sMsg"].(string) + `"}]}`
					I.account_info_deliver.InsertOutSideOrder([]byte(error_info), strategy_name)
				}
			}
		case <-time.After(time.Millisecond * 100):
		}
	}
}

func (I *InfoDeliver) DealRequest(new_request *deliver.LocalSubmit) (string, *global.SubmitInfo) {
	requestMap := make(map[string]bool)
	ans_sub := global.GenSubmitInfo()
	temp_list := strings.Split(new_request.Subtype, " ")
	for i := 0; i < len(temp_list); i++ {
		requestMap[temp_list[i]] = true
	}
	// 判断是否需要bar信息
	if requestMap["bar"] || requestMap["Bar"] || requestMap["BAR"] {
		// 获取bar相关信息，bar的Insid列表，bar的类型，传送bar的端口
		ans_sub.Bar.Judge = true
		ans_sub.Bar.InsList = strings.Split(new_request.BarInsid, " ")
		ans_sub.Bar.Custom_type = new_request.Barcustom
		ans_sub.Bar.Port = new_request.BarPort
	} else {
		ans_sub.Bar.Judge = false
	}
	if requestMap["tick"] || requestMap["Tick"] || requestMap["TICK"] {
		// 获取bar相关信息，tick的Insid列表，传送tick的端口
		ans_sub.Tick.Judge = true
		ans_sub.Tick.InsList = strings.Split(new_request.TickInsid, " ")
		ans_sub.Tick.Port = new_request.TickPort
	} else {
		ans_sub.Tick.Judge = false
	}
	// order/account/position
	if requestMap["order"] || requestMap["account"] || requestMap["position"] {
		ans_sub.Account.Judge = true
		ans_sub.Account.Port = new_request.AccountPort
		ans_sub.Account.Simulate = I.simulate
		ans_sub.Account.Userconf = I.userconfig
		ans_sub.Account.OrderJudge, ans_sub.Account.AccountJudge, ans_sub.Account.PositionJudge = requestMap["order"], requestMap["account"], requestMap["position"]
	} else {
		ans_sub.Account.Judge = false
	}
	return new_request.Strategyname, ans_sub
}

func main() {
	config := global.GetConfig("../../conf/conf.ini")
	ifd := GenInfoDeliver("6101", "6102", config.UserInfo["1"], false, 5)
	ifd.Start()
}
