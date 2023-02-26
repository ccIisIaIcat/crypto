package goClient

import (
	"log"
	"time"

	deliver "godeliver"

	"golang.org/x/net/context"

	// 导入grpc包
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"global"
)

// 用于向python端发送数据
type InfoDeliver struct {
	// pubilc
	BarChan  chan global.BarInfo
	TickChan chan global.TickInfo
	// privte
	barsignal  bool
	ticksignal bool
	barport    string
	tickport   string
}

func GenInfoDeliver() *InfoDeliver {
	ifd := &InfoDeliver{}
	ifd.barsignal = false
	ifd.ticksignal = false
	return ifd
}

func (I *InfoDeliver) ConnectBar(barchan chan global.BarInfo, port string) {
	I.barsignal = true
	I.BarChan = barchan
	I.barport = port
}

func (I *InfoDeliver) ConnectTick(tickchan chan global.TickInfo, port string) {
	I.ticksignal = true
	I.TickChan = tickchan
	I.tickport = port
}

func (I *InfoDeliver) startbar() {
	conn, err := grpc.Dial("localhost:"+I.barport, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()
	// 初始化BarDataReceiver服务客户端
	c := deliver.NewBarDataReceiverClient(conn)
	// 初始化上下文，设置请求超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// 延迟关闭请求会话
	defer cancel()
	for {
		barinfo := <-I.BarChan
		// 调用BarDataReceiver接口，发送条消息
		response, err := c.BarDataReceiver(ctx, I.CopyBar(barinfo))
		if err != nil {
			log.Println(err)
		}
		log.Println(response.ResponseMe)
	}

}

func (I *InfoDeliver) starttick() {
	conn, err := grpc.Dial("localhost:"+I.tickport, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()
	// 初始化BarDataReceiver服务客户端
	c := deliver.NewTickDataReceiverClient(conn)
	// 初始化上下文，设置请求超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// 延迟关闭请求会话
	defer cancel()
	for {
		tickinfo := <-I.TickChan
		// 调用BarDataReceiver接口，发送条消息
		response, err := c.TickDataReceiver(ctx, I.CopyTick(tickinfo))
		if err != nil {
			log.Println(err)
		}
		log.Println(response.ResponseMe)
	}

}

func (I *InfoDeliver) Start() {
	if I.barsignal {
		go I.startbar()
	}
	if I.ticksignal {
		go I.starttick()
	}
}

func (I *InfoDeliver) CopyBar(bargobal global.BarInfo) *deliver.BarData {
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

func (I *InfoDeliver) CopyTick(tickdata global.TickInfo) *deliver.TickData {
	temp := &deliver.TickData{}
	temp.Insid = tickdata.Insid
	temp.Ts_Price = int64(tickdata.Ts_Price)
	temp.Ask1Price = float32(tickdata.Ask1_price)
	temp.Bid1Price = float32(tickdata.Bid1_price)
	temp.Ask1Volumn = float32(tickdata.Ask1_volumn)
	temp.Bid1Volumn = float32(tickdata.Bid1_volumn)
	return temp
}
