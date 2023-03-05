package main

import (
	"fmt"
	"global"
	deliver "godeliver"
	"local_deliver/deliver_server"
	"strings"
	"sync"
	"time"
)

// 尝试重构服务端代码结构，开启一个接收信号的服务端，每收到一个服务请求开启对应的服务，访问不到端口后自动断开

type InfoDeliver struct {
	submitserver_port string
	orderserver_port  string
	userconfig        global.ConfigUser
	simulate          bool
}

// 赋值端口和参数
func GenInfoDeliver(submitserver_port string, orderserver_port string, userconfig global.ConfigUser, simulate_account bool) *InfoDeliver {
	ifd := &InfoDeliver{}
	ifd.submitserver_port = submitserver_port
	ifd.orderserver_port = orderserver_port
	ifd.userconfig = userconfig
	ifd.simulate = simulate_account
	return ifd
}

// 初始化服务，直接开启Orderserver,实时接收InfoDeliver收到的请求，进行对应操作
func (I *InfoDeliver) Start() {
	orderserver := deliver_server.OrderServer{}
	go orderserver.OrderServerListen(I.orderserver_port, I.userconfig, I.simulate)
	submitserver := deliver_server.SubmitServer{}
	go submitserver.SubmitServerListen(I.submitserver_port)
	fmt.Println("server starting")
	for {
		select {
		case temp := <-submitserver.InfoChan:
			fmt.Println(temp)
			if temp.Subtype != "ping" {
				go I.DealRequest(temp)
			}
		case <-time.After(time.Second):
		}
	}

}

func (I *InfoDeliver) DealRequest(new_request *deliver.LocalSubmit) {
	requestMap := make(map[string]bool)
	temp_list := strings.Split(new_request.Subtype, " ")
	for i := 0; i < len(temp_list); i++ {
		requestMap[temp_list[i]] = true
	}
	l := sync.Mutex{}
	judge_count := 0
	// 判断是否需要bar信息
	if requestMap["bar"] || requestMap["Bar"] || requestMap["BAR"] {
		// 获取bar相关信息，bar的Insid列表，bar的类型，传送bar的端口
		bar_insid_list := strings.Split(new_request.BarInsid, " ")
		bar_type := new_request.Barcustom
		bar_port := new_request.BarPort
		go func() {
			deliver_server.DeliverBar(bar_insid_list, bar_port, bar_type)
			l.Lock()
			judge_count += 1
			l.Unlock()
		}()
	} else {
		l.Lock()
		judge_count += 1
		l.Unlock()
	}
	if requestMap["tick"] || requestMap["Tick"] || requestMap["TICK"] {
		// 获取bar相关信息，tick的Insid列表，传送tick的端口
		tick_insid_list := strings.Split(new_request.TickInsid, " ")
		tick_port := new_request.TickPort
		go func() {
			deliver_server.DeliverTick(tick_insid_list, tick_port)
			l.Lock()
			judge_count += 1
			l.Unlock()
		}()
	} else {
		l.Lock()
		judge_count += 1
		l.Unlock()
	}
	// order/account/position
	if requestMap["order"] || requestMap["account"] || requestMap["position"] {
		order_sub, account_sub, position_sub := requestMap["order"], requestMap["account"], requestMap["position"]
		account_port := new_request.AccountPort
		go func() {
			deliver_server.DeliverAccount(I.userconfig, account_port, account_sub, order_sub, position_sub, I.simulate)
			l.Lock()
			judge_count += 1
			l.Unlock()
		}()
	} else {
		l.Lock()
		judge_count += 1
		l.Unlock()
	}
	for {
		time.Sleep(time.Second)
		l.Lock()
		if judge_count == 3 {
			break
		}
		l.Unlock()
	}

}

func main() {
	config := global.GetConfig("../../conf/conf.ini")
	ifd := GenInfoDeliver("6101", "6102", config.UserInfo["Simulate"], true)
	ifd.Start()
}
