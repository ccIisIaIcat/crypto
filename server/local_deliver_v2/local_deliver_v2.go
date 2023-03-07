package main

import (
	"fmt"
	"global"
	"local_deliver_v2/deliver_server"
	"time"
)

// 尝试重构服务端代码结构，开启一个接收信号的服务端，每收到一个服务请求开启对应的服务，访问不到端口后自动断开
// v2更新：1、增加ping pong机制，当一定时间未收到客户端的ping时断开连接
//
//	2、对于报单，先放在服务端缓存内防止客户端卡顿，在报单place完毕后，若开启order频道，通过order deliver返回客户端
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

func Start_Submit_Server() {

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
			go I.DealRequest(temp)
		case <-time.After(time.Second):
		}
	}

}
