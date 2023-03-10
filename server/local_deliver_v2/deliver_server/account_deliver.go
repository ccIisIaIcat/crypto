package deliver_server

import (
	"account"
	"context"
	"encoding/json"
	"global"
	deliver "godeliver"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// account_deliver改为收到信息向对应的策略端口发送信息
type Account_deliver struct {
	ac                 *account.Account
	Port_map           map[string]string // 每一个策略对应的端口
	account_info_chan  map[string](chan []byte)
	account_subscribe  map[string]bool
	position_subscribe map[string]bool
	order_subsicribe   map[string]bool
	PingPongMapChan    map[string](chan bool)
	timeout            int
	// local_lock
	lock sync.Mutex
	// local_log
	order_log    *log.Logger
	position_log *log.Logger
}

func GenAccountDeliver(userconf global.ConfigUser, simulate_account bool, timeout int) *Account_deliver {
	acd := &Account_deliver{}
	acd.timeout = timeout
	acd.ac = account.GenAccount(userconf, true, true, true, simulate_account)
	acd.Port_map = make(map[string]string)
	acd.account_info_chan = make(map[string]chan []byte)
	acd.account_subscribe = make(map[string]bool)
	acd.position_subscribe = make(map[string]bool)
	acd.order_subsicribe = make(map[string]bool)
	acd.PingPongMapChan = make(map[string]chan bool)
	acd.lock = sync.Mutex{}
	// 本地账户日志
	// 订单日志
	file, err := os.OpenFile("./log/order.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}
	acd.order_log = log.New(io.MultiWriter(file, os.Stderr), "[order]:", log.Ldate|log.Ltime|log.Lshortfile)
	// 仓位日志
	file2, err := os.OpenFile("./log/position.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}
	acd.position_log = log.New(io.MultiWriter(file2, os.Stderr), "position", log.Ldate|log.Ltime|log.Lshortfile)
	return acd
}

func (A *Account_deliver) DeliverAccount() {
	go A.ac.Start()
	time.Sleep(time.Second)
	go A.deliverByStrategyName()
	global.NeverStop()
}

func (A *Account_deliver) AddStrategy(strategy_name string, port string, account_sub bool, position_sub bool, order_sub bool) {
	A.lock.Lock()
	A.account_info_chan[strategy_name] = make(chan []byte, 100)
	A.Port_map[strategy_name] = port
	if account_sub {
		A.account_subscribe[strategy_name] = true
	}
	if position_sub {
		A.position_subscribe[strategy_name] = true
	}
	if order_sub {
		A.order_subsicribe[strategy_name] = true
	}
	A.PingPongMapChan[strategy_name] = make(chan bool, 100)
	A.lock.Unlock()
	go A.pingpongCheck(A.timeout, strategy_name)
	A.startaccount(strategy_name)
}

func (A *Account_deliver) CancelStrategy(strategy_name string) {
	A.lock.Lock()
	if _, ok := A.Port_map[strategy_name]; ok {
		log.Println("delet ", strategy_name, " from port map")
		delete(A.Port_map, strategy_name)
	}
	if _, ok := A.account_subscribe[strategy_name]; ok {
		log.Println("delet ", strategy_name, " from account_subscribe")
		delete(A.account_subscribe, strategy_name)
	}
	if _, ok := A.position_subscribe[strategy_name]; ok {
		log.Println("delet ", strategy_name, " from position_subscribe")
		delete(A.position_subscribe, strategy_name)
	}
	if _, ok := A.order_subsicribe[strategy_name]; ok {
		log.Println("delet ", strategy_name, " from order_subsicribe")
		delete(A.order_subsicribe, strategy_name)
	}
	if _, ok := A.account_info_chan[strategy_name]; ok {
		log.Println("delet ", strategy_name, " from account_info_chan")
		delete(A.account_info_chan, strategy_name)
	}
	if _, ok := A.account_info_chan[strategy_name]; ok {
		log.Println("delet ", strategy_name, " from pingpong_chan")
		delete(A.PingPongMapChan, strategy_name)
	}
	A.lock.Unlock()
}

func (A *Account_deliver) deliverByStrategyName() {
	for {
		select {
		case info := <-A.ac.InfoChanAccount:
			A.lock.Lock()
			for k := range A.account_subscribe {
				A.account_info_chan[k] <- info
			}
			A.lock.Unlock()
		case info := <-A.ac.InfoChanPositions:
			// A.position_log.Println("------------------------------position-----------------------------")
			// A.position_log.Println(string(info))
			A.lock.Lock()
			for k := range A.position_subscribe {
				A.account_info_chan[k] <- info
			}
			A.lock.Unlock()
		case info := <-A.ac.InfoChanOrders:
			A.order_log.Println(string(info))
			A.lock.Lock()
			var temp_judge map[string]interface{}
			json.Unmarshal(info, &temp_judge)
			strateName := strings.Split(temp_judge["data"].([]interface{})[0].(map[string]interface{})["clOrdId"].(string), "0")[0]
			if _, ok := A.order_subsicribe[strateName]; ok {
				A.account_info_chan[strateName] <- info
			}
			A.lock.Unlock()
		case <-time.After(time.Millisecond * 10):
		}
	}
}

func (A *Account_deliver) InsertOutSideOrder(info []byte, strategy_name string) {
	A.account_info_chan[strategy_name] <- info
}

func (A *Account_deliver) startaccount(strategyName string) {
	conn, err := grpc.Dial("localhost:"+A.Port_map[strategyName], grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	signal := false
	defer cancel()
	for {
		select {
		case json_info := <-A.account_info_chan[strategyName]:
			response, err := c.JsonReceiver(ctx, A.copy_json(json_info))
			if err != nil {
				log.Println(err)
				signal = true
			}
			log.Println(response)
		case <-time.After(time.Millisecond * 20):
		}
		A.lock.Lock()
		if _, ok := A.Port_map[strategyName]; !ok {
			A.lock.Unlock()
			break
		}
		if signal {
			A.lock.Unlock()
			break
		}
		A.lock.Unlock()
		// fmt.Println(barinfo)
		// 调用BarDataReceiver接口，发送条消息

	}
}

func (A *Account_deliver) pingpongCheck(time_out int, strategy_name string) {
	temp_signal := true
	for temp_signal {
		select {
		case <-A.PingPongMapChan[strategy_name]:
			// log.Println("ping >>>>>>>")
		case <-time.After(time.Second * time.Duration(time_out)):
			A.CancelStrategy(strategy_name)
			log.Println("strategy:", strategy_name, "account pingpong timeout! disconnected")
			temp_signal = false
		}
	}
}

func (A *Account_deliver) copy_json(temp_json []byte) *deliver.JsonInfo {
	temp := &deliver.JsonInfo{}
	temp.Jsoninfo = string(temp_json)
	return temp

}
