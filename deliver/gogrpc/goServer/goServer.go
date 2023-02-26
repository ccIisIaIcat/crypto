package main

import (
	"context"
	"fmt"
	deliver "godeliver"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	deliver.UnimplementedOrerReceiverServer
}

func (S *Server) OrerRReceiver(ctx context.Context, in *deliver.Order) (*deliver.Response, error) {
	fmt.Println(in)
	temp := &deliver.Response{ResponseMe: "Get"}

	return temp, nil
}

func main() {
	// 监听127.0.0.1:50051地址
	lis, err := net.Listen("tcp", "127.0.0.1:4352")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc服务端
	s := grpc.NewServer()

	// 注册Greeter服务
	deliver.RegisterOrerReceiverServer(s, &Server{})

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
