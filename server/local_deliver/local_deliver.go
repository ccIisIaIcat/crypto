package main

import (
	"fmt"
	"global"
	"goClient"
	"query_tick"
	"time"
)

func main() {
	fmt.Println("lalala")
	tick_chan := make(chan *global.TickInfo, 200)
	qt := query_tick.QueryTick{InsId_list: []string{"BTC-USDT-SWAP"}, Tick_info_chan: tick_chan}
	go qt.Start()
	time.Sleep(time.Second)

	gfd := goClient.GenInfoDeliver()
	gfd.ConnectTick(qt.Tick_info_chan, "3903")
	gfd.Start()

	global.NeverStop()

}
