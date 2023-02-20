package main

import (
	"fmt"
	"log"
	"query_detail"
	"query_insid"
	"time"
)

func main() {
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
