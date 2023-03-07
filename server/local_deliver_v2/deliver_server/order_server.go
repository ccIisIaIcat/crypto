package deliver_server

import (
	"context"
	"fmt"
	"global"
	deliver "godeliver"
	"log"
	"net"
	"reflect"
	"strings"
	"time"
	"trade_restful"
	"trade_restful_simulate"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type OrderServer struct {
	// public
	Res_chan chan string // 对外接口，用于判断
	// private
	deliver.UnimplementedOrerReceiverServer
	trade_server          *trade_restful.TradeRestful
	trade_server_simulate *trade_restful_simulate.TradeRestfulSimulate
	userconfig            global.ConfigUser
	simulate              bool
	order_chan            chan *deliver.Order
}

// 对receiver模块进行更新，当收到报单后，直接存放在order_chan中并迅速给出回执，再启动一个服务专门处理订单，orderplaced信息对外推送
func (O *OrderServer) OrerRReceiver(ctx context.Context, order *deliver.Order) (*deliver.Response, error) {
	O.order_chan <- order
	return &deliver.Response{ResponseMe: "get order"}, nil
}

func (O *OrderServer) sendOrder() {
	for {
		select {
		case temp_order := <-O.order_chan:
			format_order := genorder(temp_order)
			if O.simulate {
				res := O.trade_server_simulate.SendOrder(format_order)
				O.Res_chan <- res
			} else {
				res := O.trade_server.SendOrder(format_order)
				O.Res_chan <- res
			}
		case <-time.After(time.Millisecond * 50):
		}
	}

}

func (O *OrderServer) OrderServerListen(port string, userconf global.ConfigUser, simulate bool) {
	O.simulate = simulate
	O.userconfig = userconf
	O.Res_chan = make(chan string, 100)
	O.order_chan = make(chan *deliver.Order, 100)
	if simulate {
		O.trade_server_simulate = trade_restful_simulate.GenTradeRestfulSimulate(userconf)
	} else {
		O.trade_server = trade_restful.GenTradeRestful(userconf)
	}
	lis, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go O.sendOrder()
	// 实例化grpc服务端
	s := grpc.NewServer()

	// 注册Greeter服务
	deliver.RegisterOrerReceiverServer(s, O)

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func genorder(od *deliver.Order) string {
	fmt.Println(od)
	answer := `{"instId":"` + od.InsId + `"`
	// temp_list := []string{"insId", "tdMode", "ccy", "clOrdId", "tag", "side", "posSide", "ordType", "sz", "px", "reduceOnly", "tgtCcy", "banAmend", "tpTriggerPx", "tpOrdPx", "slTriggerPx", "slOrdPx", "tpTriggerPxType", "slTriggerPxType", "quickMgnType", "brokerID"}
	hofvalue := reflect.ValueOf(*od)
	tp := reflect.TypeOf(*od)
	for i := 0; i < tp.NumField(); i++ {
		temp_name := strings.ToLower(string(tp.Field(i).Name[0])) + string(tp.Field(i).Name[1:])
		if tp.Field(i).Name == "brokerID" {
			continue
		} else if tp.Field(i).Name == "reduceOnly" || tp.Field(i).Name == "banAmend" {
			answer += `,"` + tp.Field(i).Name + `":` + hofvalue.Field(i).Interface().(string) + ``
			continue
		} else if temp_name == "tdMode" || temp_name == "ccy" || temp_name == "clOrdId" || temp_name == "tag" || temp_name == "side" || temp_name == "posSide" || temp_name == "ordType" || temp_name == "sz" || temp_name == "px" || temp_name == "tgtCcy" || temp_name == "tpTriggerPx" || temp_name == "tpOrdPx" || temp_name == "slTriggerPx" || temp_name == "slOrdPx" || temp_name == "tpTriggerPxType" || temp_name == "slTriggerPxType" || temp_name == "quickMgnType" {
			if hofvalue.Field(i).Interface().(string) != "" {
				answer += `,"` + temp_name + `":"` + hofvalue.Field(i).Interface().(string) + `"`
			}
		}
	}
	answer += `}`
	fmt.Println(answer)
	return answer
}
