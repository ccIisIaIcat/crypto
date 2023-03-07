package deliver_server

import (
	"context"
	deliver "godeliver"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 用于接收请求,向chan中传送请求
type SubmitServer struct {
	deliver.UnimplementedSubmitServerReceiverServer
	InfoChan chan *deliver.LocalSubmit
}

func (S *SubmitServer) SubmitServerReceiver(ctx context.Context, submit *deliver.LocalSubmit) (*deliver.Response, error) {
	S.InfoChan <- submit
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