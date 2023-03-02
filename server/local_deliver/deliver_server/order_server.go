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
	"trade_restful"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type OrderServer struct {
	deliver.UnimplementedOrerReceiverServer
	trade_server *trade_restful.TradeRestful
	userconfig   global.ConfigUser
}

func (O *OrderServer) OrerRReceiver(ctx context.Context, order *deliver.Order) (*deliver.Response, error) {
	format_order := genorder(order)
	res := O.trade_server.SendOrder(format_order)
	temp := &deliver.Response{ResponseMe: res}
	fmt.Println(res)
	return temp, nil
}

func (O *OrderServer) OrderServerListen(port string, userconf global.ConfigUser) {
	O.userconfig = userconf
	O.trade_server = trade_restful.GenTradeRestful(userconf)
	lis, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
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
