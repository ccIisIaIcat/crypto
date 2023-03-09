package deliver_server

import (
	"account"
	"context"
	"encoding/json"
	"global"
	deliver "godeliver"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 用于接收请求,向chan中传送请求
type SubmitServer struct {
	deliver.UnimplementedSubmitServerReceiverServer
	InfoChan chan *deliver.LocalSubmit
	Userconf global.ConfigUser
	Simulate bool
}

func (S *SubmitServer) SubmitServerReceiver(ctx context.Context, submit *deliver.LocalSubmit) (*deliver.Response, error) {
	S.InfoChan <- submit
	// 后续在此处添加对api初始化的回执
	if submit.Initjson != "" {
		temp_acc := account.GenAccountConf(S.Userconf, S.Simulate)
		var temp map[string]interface{}
		json.Unmarshal([]byte(submit.Initjson), &temp)
		temp_acc.SetPositionMode(temp["TradingMode"].(string))
		for k, v := range temp["LeverageSet"].(map[string]interface{}) {
			temp_list := strings.Split(k, " ")
			if len(temp_list) == 2 {
				temp_acc.SetLeverage(temp_list[0], v.(string), temp_list[1])
			} else {
				temp_acc.SetLeverage(temp_list[0], v.(string), "cross")
			}
		}
		if temp["TradingInsid"].(string) != "" {
			temp_list := strings.Split(temp["TradingInsid"].(string), " ")
			if len(strings.Split(temp_list[0], "-")) == 2 {
				respon_json := temp_acc.GetInsIdInfo("SWAP", strings.Split(temp["TradingInsid"].(string), " "))
				temp := &deliver.Response{ResponseMe: string(respon_json)}
				return temp, nil
			} else {
				if strings.Split(temp_list[0], "-")[2] == "SWAP" {
					respon_json := temp_acc.GetInsIdInfo("SWAP", strings.Split(temp["TradingInsid"].(string), " "))
					temp := &deliver.Response{ResponseMe: string(respon_json)}
					return temp, nil
				} else {
					respon_json := temp_acc.GetInsIdInfo("FUTURES", strings.Split(temp["TradingInsid"].(string), " "))
					temp := &deliver.Response{ResponseMe: string(respon_json)}
					return temp, nil
				}
			}
		}

	}
	temp := &deliver.Response{ResponseMe: "Get"}
	return temp, nil
}

func (S *SubmitServer) SubmitServerListen(port string) {
	S.InfoChan = make(chan *deliver.LocalSubmit, 100)
	lis, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化grpc服务端
	s := grpc.NewServer()

	// 注册Greeter服务
	deliver.RegisterSubmitServerReceiverServer(s, S)

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
