package deliver_server

import (
	"context"
	"global"
	deliver "godeliver"
	"log"
	"query_bar_custom"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 用于根据指定字段访问websocket服务，获取tick数据，并发送给指定端口
type Bar_deliver struct {
	InsList   []string
	query_bar query_bar_custom.QueryBar
	Signal    bool
}

func (B *Bar_deliver) DeliverBar(Ins_list []string, Port string, Custom_type string) int {
	go B.query_bar.Start()
	time.Sleep(time.Second)
	go B.startbarcustom(Port)
	for !B.Signal {
		time.Sleep(time.Second)
	}
	B.query_bar.Close()
	return 1
}

func GenBarDeliver(Ins_list []string, Port string, Custom_type string) *Bar_deliver {
	bd := &Bar_deliver{}
	bd.query_bar = query_bar_custom.QueryBar{InsId_list: Ins_list, Custom_type: Custom_type}
	bd.Signal = false
	return bd
}

func (B *Bar_deliver) startbarcustom(Port string) {
	conn, err := grpc.Dial("localhost:"+Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()
	// 初始化BarDataReceiver服务客户端
	c := deliver.NewCustomDataReceiverClient(conn)
	// 初始化上下文，设置请求超时时间
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	ctx, cancel := context.WithCancel(context.Background())
	// 延迟关闭请求会话
	defer cancel()
	for {
		barinfo := <-B.query_bar.Bar_info_chan
		// fmt.Println(barinfo)
		// 调用BarDataReceiver接口，发送条消息
		_, err := c.CustomDataReceiver(ctx, B.CopyBar(barinfo))
		if err != nil {
			log.Println(err)
			B.Signal = true
		}
		// log.Println(response)
	}
}

func (B *Bar_deliver) CopyBar(bargobal *global.BarInfo) *deliver.BarData {
	temp := &deliver.BarData{}
	temp.Insid = bargobal.Insid
	temp.TsOpen = int64(bargobal.Ts_open)
	temp.OpenPrice = float32(bargobal.Open_price)
	temp.HighPrice = float32(bargobal.Open_price)
	temp.LowPrice = float32(bargobal.Low_price)
	temp.ClosePrice = float32(bargobal.Close_price)
	temp.Vol = float32(bargobal.Vol)
	temp.VolCcy = float32(bargobal.VolCcy)
	temp.VolCcyQuote = float32(bargobal.VolCcyQuote)
	temp.Oi = float32(bargobal.Oi)
	temp.OiCcy = float32(bargobal.OiCcy)
	temp.TsOi = int64(bargobal.Ts_oi)
	temp.FundingRate = float32(bargobal.FundingRate)
	temp.NextFundingRate = float32(bargobal.NextFundingRate)
	temp.TS_NextFundingRate = int64(bargobal.TS_NextFundingRate)
	temp.Ts_FundingRate = int64(bargobal.Ts_FundingRate)
	return temp
}
