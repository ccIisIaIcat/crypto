package deliver_server

import (
	"account"
	"context"
	"global"
	deliver "godeliver"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Account_deliver struct {
	ac     *account.Account
	Signal bool
}

func DeliverAccount(userconf global.ConfigUser, Port string, account_sub bool, order_sub bool, position_sub bool, simulate_account bool) int {
	acd := genAccountDeliver(userconf, account_sub, order_sub, position_sub, simulate_account)
	go acd.ac.Start()
	time.Sleep(time.Second)
	go acd.startaccount(Port)
	for !acd.Signal {
		time.Sleep(time.Second)
	}
	acd.ac.Close()
	return 1
}
func genAccountDeliver(userconf global.ConfigUser, account_sub bool, order_sub bool, position_sub bool, simulate_account bool) *Account_deliver {
	acd := &Account_deliver{}
	acd.ac = account.GenAccount(userconf, account_sub, order_sub, position_sub, simulate_account)
	acd.Signal = false

	return acd
}
func (A *Account_deliver) startaccount(Port string) {
	conn, err := grpc.Dial("localhost:"+Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()
	// 初始化BarDataReceiver服务客户端
	c := deliver.NewJsonReceiverClient(conn)
	// 初始化上下文，设置请求超时时间
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	ctx, cancel := context.WithCancel(context.Background())
	// 延迟关闭请求会话
	defer cancel()
	for {
		select {
		case json_info := <-A.ac.InfoChanAccount:
			response, err := c.JsonReceiver(ctx, A.copy_json(json_info))
			if err != nil {
				log.Println(err)
				A.Signal = true
			}
			log.Println(response)
		case json_info := <-A.ac.InfoChanOrders:
			response, err := c.JsonReceiver(ctx, A.copy_json(json_info))
			if err != nil {
				log.Println(err)
				A.Signal = true
			}
			log.Println(response)
		case json_info := <-A.ac.InfoChanPositions:
			response, err := c.JsonReceiver(ctx, A.copy_json(json_info))
			if err != nil {
				log.Println(err)
				A.Signal = true
			}
			log.Println(response)
		}
		// fmt.Println(barinfo)
		// 调用BarDataReceiver接口，发送条消息

	}
}

func (A *Account_deliver) copy_json(temp_json []byte) *deliver.JsonInfo {
	temp := &deliver.JsonInfo{}
	temp.Jsoninfo = string(temp_json)
	return temp

}
