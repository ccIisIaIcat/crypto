package main

import (
	"context"
	"deliver"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server struct {
	deliver.UnimplementedBarDataRevicerServer
}

func (S *Server) BarDataRevicer(ctx context.Context, in *deliver.BarData) (*deliver.Response, error) {
	fmt.Println(in)
	return nil, status.Errorf(codes.Unimplemented, "method BarDataRevicer not implemented")
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
	deliver.RegisterBarDataRevicerServer(s, &Server{})

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
