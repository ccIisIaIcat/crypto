package main

import (
	"fmt"
	"global"
	"time"
	"websocketlocal"
)

type Account struct {
	// public
	InfoChan chan []byte
	// private
	ws       *websocketlocal.WebSocketLocal
	userconf global.ConfigUser
}

func GenAccount(userconf global.ConfigUser) {
	ac := Account{}
	ac.ws = websocketlocal.GenWebSocket("wss://ws.okx.com:8443/ws/v5/private", 10)
	ac.userconf = userconf
	go ac.ws.StartGather()

	global.NeverStop()

}

func main() {
	fmt.Println(time.Now())
}
