package goClient

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	deliver "godeliver"

	"golang.org/x/net/context"

	// 导入grpc包
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"global"
)

// intro:
// 1、先开启一个go端服务，收到python端报单请求后，实时fasongtick和bar数据

// 用于向python端发送数据
type InfoDeliver struct {
	// pubilc
	TickChan      chan *global.TickInfo
	BarCustomChan chan *global.BarInfo
	// privte
	ticksignal      bool
	barcustomsignal bool
	tickport        string
	barcustomport   string
	Local_server    *Server
	Stop_signal     sync.Map
}

type Server struct {
	deliver.UnimplementedOrerReceiverServer
	InfoChan chan *deliver.Order
}

func (S *Server) OrerRReceiver(ctx context.Context, in *deliver.Order) (*deliver.Response, error) {
	fmt.Println(in)
	S.InfoChan <- in
	temp := &deliver.Response{ResponseMe: "Get"}

	return temp, nil
}

func GenInfoDeliver() *InfoDeliver {
	ifd := &InfoDeliver{}
	ifd.ticksignal = false
	ifd.barcustomsignal = false
	ifd.Local_server = &Server{}
	ifd.Local_server.InfoChan = make(chan *deliver.Order, 100)
	ifd.Stop_signal = sync.Map{}
	ifd.Stop_signal.Store("disconnect", false)
	return ifd
}

func (I *InfoDeliver) ConnectTick(tickchan chan *global.TickInfo, port string) {
	I.ticksignal = true
	I.TickChan = tickchan
	I.tickport = port
}

func (I *InfoDeliver) ConnectBarCustom(barchan chan *global.BarInfo, port string) {
	I.barcustomsignal = true
	I.BarCustomChan = barchan
	I.barcustomport = port
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
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	ctx, cancel := context.WithCancel(context.Background())
	// 延迟关闭请求会话
	defer cancel()
	for {
		tickinfo := <-I.TickChan
		fmt.Println(tickinfo)
		// 调用BarDataReceiver接口，发送条消息
		response, err := c.TickDataReceiver(ctx, I.CopyTick(tickinfo))
		if err != nil {
			log.Println(err)
			I.Stop_signal.Store("disconnect", true)
			break
		}
		log.Println(response)
	}
}

func (I *InfoDeliver) startbarcustom() {
	conn, err := grpc.Dial("localhost:"+I.barcustomport, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		barinfo := <-I.BarCustomChan
		fmt.Println(barinfo)
		// 调用BarDataReceiver接口，发送条消息
		response, err := c.CustomDataReceiver(ctx, I.CopyBar(barinfo))
		if err != nil {
			log.Println(err)
			I.Stop_signal.Store("disconnect", true)
			break
		}
		log.Println(response)
	}
}

func (I *InfoDeliver) LocalServerListen(port string) {
	lis, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化grpc服务端
	s := grpc.NewServer()

	// 注册Greeter服务
	deliver.RegisterOrerReceiverServer(s, I.Local_server)

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (I *InfoDeliver) StartDeliver() {
	//
	if I.ticksignal {
		go I.starttick()
	}
	if I.barcustomsignal {
		go I.startbarcustom()
	}

	judge, _ := I.Stop_signal.Load("disconnect")
	fmt.Println(judge)
	for !judge.(bool) {
		time.Sleep(time.Second)
		judge, _ = I.Stop_signal.Load("disconnect")
	}

}

func (I *InfoDeliver) CopyBar(bargobal *global.BarInfo) *deliver.BarData {
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

func (I *InfoDeliver) CopyTick(tickdata *global.TickInfo) *deliver.TickData {
	temp := &deliver.TickData{}
	temp.Insid = tickdata.Insid
	temp.Ts_Price = int64(tickdata.Ts_Price)
	temp.Ask1Price = float32(tickdata.Ask1_price)
	temp.Bid1Price = float32(tickdata.Bid1_price)
	temp.Ask1Volumn = float32(tickdata.Ask1_volumn)
	temp.Bid1Volumn = float32(tickdata.Bid1_volumn)
	return temp
}
