package main

import (
	"fmt"
	"log"
	"query_detail"
	"query_insid"
	"time"
)

func main() {
	// a := websocketlocal.GenWebSocket("wss://ws.okx.com:8443/ws/v5/public", 10)
	// sample_sub := `{"op": "subscribe","args": [{"channel": "tickers","instId": "BTC-USDT-SWAP"}]}`

	// a.Submit([]byte(sample_sub), true)
	// time.Sleep(time.Second * 2)
	// go a.StartGather()
	// time.Sleep(time.Second * 3)
	// a.Close()
	// global.NeverStop()
	fmt.Println(query_insid.GetInsByType("SWAP"))
	a := query_detail.QueryDetail{InsId_list: query_insid.GetInsByType("SWAP")}
	go a.Start()
	// fmt.Println("lala")
	time.Sleep(time.Second)
	for {
		temp := <-a.Tick_info_chan
		log.Println(temp)
	}

}
