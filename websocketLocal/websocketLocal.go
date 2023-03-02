package websocketlocal

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 用于建立websocket连接，发送订阅请求，提供回执响应接口
type WebSocketLocal struct {
	// Public
	Address     string      // websocket 连接地址
	TimeCounter int         // 计时器响应时长，用于断线重连
	InfoChan    chan []byte // 对外的消息接收通道

	// Private
	conn         *websocket.Conn // websocket 连接对象
	subcribeInfo []string        // 正在订阅的服务，用于断线重连恢复
	temp_chan    chan []byte
	check_map    sync.Map
	signal       bool
}

// 生成一个websocket对象
func GenWebSocket(address string, timecounter int) *WebSocketLocal {
	ws := WebSocketLocal{}
	ws.Address = address
	ws.TimeCounter = timecounter
	ws.signal = true
	ws.subcribeInfo = make([]string, 0)
	ws.InfoChan = make(chan []byte, 100)
	ws.temp_chan = make(chan []byte, 1000)
	ws.check_map = sync.Map{}
	ws.check_map.Store("check", true)
	// 尝试握手(若失败，重复，重复100次失败报错)
	dialer := websocket.Dialer{}
	var err error
	for i := 0; i < 100; i++ {
		ws.conn, _, err = dialer.Dial(address, nil)
		if err != nil {
			log.Println("websocket handle failed, try reconnect:", i)
		} else {
			log.Println("websocket handle successed!")
			break
		}
	}
	if err != nil {
		panic(err)
	}
	return &ws
}

// 提交订阅或取消订阅
func (W *WebSocketLocal) Submit(SubInfo []byte, save bool) {
	if string(SubInfo) != "ping" && save {
		W.subcribeInfo = append(W.subcribeInfo, string(SubInfo))
	}
	err := W.conn.WriteMessage(websocket.TextMessage, SubInfo)
	if err != nil {
		log.Println("submit failed:", err)
	} else {
		log.Println("submited")
	}

}

// 一个用于预处理的私有函数
func (W *WebSocketLocal) preprocess() {
	for W.signal {
		check, _ := W.check_map.Load("check")
		for !check.(bool) {
			time.Sleep(time.Second)
			fmt.Println(time.Now())
			check, _ = W.check_map.Load("check")
		}
		messageType, info, err := W.conn.ReadMessage()
		if err != nil || messageType == -1 {
			if W.signal {
				log.Println("err while get respon", err)
				W.restartWebsocket()
			}
		} else {
			W.temp_chan <- info
		}
	}
}

// 开启获取服务
func (W *WebSocketLocal) StartGather() {
	go W.preprocess()
	for W.signal {
		check, _ := W.check_map.Load("check")
		for !check.(bool) {
			time.Sleep(time.Second)
			fmt.Println(time.Now())
			check, _ = W.check_map.Load("check")
		}
		select {
		case <-time.After(5 * time.Second):
			W.Submit([]byte("ping"), false)
			select {
			case <-time.After(10 * time.Second):
				if W.signal {
					log.Println("waiting ping pong timeout, restart websocket")
					W.restartWebsocket()
				}
			case temp := <-W.temp_chan:
				if string(temp) != "pong" {
					W.InfoChan <- temp
				}
			}

		case temp := <-W.temp_chan:
			if string(temp) != "pong" {
				W.InfoChan <- temp
			}
		}
	}

}

// 一个用于断线重连的私有函数
func (W *WebSocketLocal) restartWebsocket() {
	W.check_map.Store("check", false)
	log.Println("restarting")
	// 尝试握手(若失败，重复,间隔3秒，重复400次失败报错)
	dialer := websocket.Dialer{}
	var err error
	for i := 0; i < 400; i++ {
		W.conn, _, err = dialer.Dial(W.Address, nil)
		if err != nil {
			log.Println("websocket handle failed, try reconnect:", i)
		} else {
			log.Println("websocket handle successed!")
			break
		}
		time.Sleep(time.Second * 5)
	}
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(W.subcribeInfo); i++ {
		W.Submit([]byte(W.subcribeInfo[i]), false)
	}
	W.check_map.Store("check", true)

}

func (W *WebSocketLocal) Close() {
	if W.conn != nil {
		W.signal = false
		W.conn.Close()
		fmt.Println("close")
	}

}

// func main() {
// 	a := GenWebSocket("wss://ws.okx.com:8443/ws/v5/public", 10)
// 	sample_sub := `{"op": "subscribe","args": [{"channel": "tickers","instId": "BTC-USDT-SWAP"}]}`

// 	a.Submit([]byte(sample_sub), true)
// 	time.Sleep(time.Second * 2)
// 	go a.StartGather()
// 	time.Sleep(time.Second * 3)
// 	a.conn.Close()
// 	NeverStop()
// }
