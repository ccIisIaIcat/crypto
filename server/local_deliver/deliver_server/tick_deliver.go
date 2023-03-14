package deliver_server

import (
	"context"
	"global"
	deliver "godeliver"
	"log"
	"query_tick"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 用于根据指定字段访问websocket服务，获取tick数据，并发送给指定端口
type Tick_deliver struct {
	InsList    []string
	query_tick query_tick.QueryTick
	Signal     bool
}

func DeliverTick(Ins_list []string, Port string) int {
	td := genTickDeliver(Ins_list, Port)
	go td.query_tick.Start()
	time.Sleep(time.Second)
	go td.starttick(Port)
	for !td.Signal {
		time.Sleep(time.Second)
	}
	td.query_tick.Close()
	// global.NeverStop()
	return 1
}

func genTickDeliver(Ins_list []string, Port string, simulate bool) *Tick_deliver {
	td := &Tick_deliver{}
	td.query_tick = query_tick.QueryTick{InsId_list: Ins_list, Simulate: simulate}
	td.Signal = false
	return td
}

func (T *Tick_deliver) starttick(tickport string) {
	conn, err := grpc.Dial("localhost:"+tickport, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := deliver.NewTickDataReceiverClient(conn)
	// 初始化上下文，设置请求超时时间
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		tickinfo := <-T.query_tick.Tick_info_chan
		// fmt.Println(tickinfo)
		// 调用BarDataReceiver接口，发送条消息
		_, err := c.TickDataReceiver(ctx, T.CopyTick(tickinfo))
		if err != nil {
			log.Println(err)
			T.Signal = true
		}
		// log.Println(response)
	}
}

func (T *Tick_deliver) CopyTick(tickdata *global.TickInfo) *deliver.TickData {
	temp := &deliver.TickData{}
	temp.Insid = tickdata.Insid
	temp.Ts_Price = int64(tickdata.Ts_Price)
	temp.Ask1Price = float32(tickdata.Ask1_price)
	temp.Bid1Price = float32(tickdata.Bid1_price)
	temp.Ask1Volumn = float32(tickdata.Ask1_volumn)
	temp.Bid1Volumn = float32(tickdata.Bid1_volumn)
	return temp
}
