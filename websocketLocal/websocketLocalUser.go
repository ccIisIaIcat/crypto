package websocketlocal

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"global"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 当遇到login信息时，更新时间戳，更新签名
// 用于建立websocket连接，发送订阅请求，提供回执响应接口
type WebSocketLocalUser struct {
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
	userconf     global.ConfigUser
}

// 生成一个websocket对象
func GenWebSocketUser(address string, timecounter int, userconf global.ConfigUser) *WebSocketLocalUser {
	ws := WebSocketLocalUser{}
	ws.userconf = userconf
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
		// ws.conn.SetWriteDeadline(time.Now().Add(30 * time.Second))
		// err := conn.WriteMessage(ws.BinaryMessage, frame)
		// if err != nil {
		// 	log.Println("write err:", err.Error())
		// 	return
		// }
		// conn.SetWriteDeadline(time.Time{})

	}
	if err != nil {
		panic(err)
	}
	return &ws
}

func ComputeHmacSha256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (W *WebSocketLocalUser) genSign(Secretkey string) (string, string) {
	method := "GET"
	requestPath := "/users/self/verify"
	temp_ts := strconv.Itoa(int(time.Now().UTC().Unix()))
	return temp_ts, ComputeHmacSha256(temp_ts+method+requestPath, Secretkey)
}

// 提交订阅或取消订阅
func (W *WebSocketLocalUser) Submit(SubInfo []byte, save bool) {
	if string(SubInfo) != "ping" && save {
		W.subcribeInfo = append(W.subcribeInfo, string(SubInfo))
	}
	err := W.conn.WriteMessage(websocket.TextMessage, SubInfo)
	if err != nil {
		log.Println("submit failed:", err)
	} else {
		if string(SubInfo) != "ping" {
			log.Println("submited:", string(SubInfo))
		}
	}

}

// 一个用于预处理的私有函数
func (W *WebSocketLocalUser) preprocess() {
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
func (W *WebSocketLocalUser) StartGather() {
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
func (W *WebSocketLocalUser) restartWebsocket() {
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
		var temp map[string]interface{}
		json.Unmarshal([]byte(W.subcribeInfo[i]), &temp)
		if temp["op"] == "login" {
			W.Submit([]byte(W.updateLogin()), false)
			time.Sleep(time.Second * 2)
		} else {
			W.Submit([]byte(W.subcribeInfo[i]), false)
		}

	}
	W.check_map.Store("check", true)

}

func (W *WebSocketLocalUser) updateLogin() string {
	temps, sign := W.genSign(W.userconf.Secretkey)
	new_login := `{"op": "login","args":[{"apiKey":"` + W.userconf.Apikey + `","passphrase" :"` + W.userconf.Passphrase + `","timestamp" :"` + temps + `","sign" :"` + sign + `" }]}`
	return new_login
}

func (W *WebSocketLocalUser) Close() {
	if W.conn != nil {
		W.signal = false
		W.conn.Close()
		fmt.Println("close")
	}
}
