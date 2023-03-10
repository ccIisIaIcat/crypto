package main

import (
	"log"
	"time"

	"deliver"

	"golang.org/x/net/context"

	// 导入grpc包
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// 导入刚才我们生成的代码所在的proto包。
)

func main() {
	// 连接grpc服务器
	conn, err := grpc.Dial("localhost:3902", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()

	// 初始化Greeter服务客户端
	c := deliver.NewBarDataReceiverClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 调用SayHello接口，发送一条消息
	a, err := c.BarDataReceiver(ctx, &deliver.BarData{Insid: "BTC-USDT", OpenPrice: 15.23})
	if err != nil {
		log.Println(err)
	}
	log.Println(a.ResponseMe)

}
