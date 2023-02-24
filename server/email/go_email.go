package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"websocketlocal"

	"gopkg.in/gomail.v2"
)

// {
//     "arg": {
//         "channel": "status"
//     },
//     "data": [
//         {
//             "begin": "1672823400000",
//             "end": "1672825980000",
//             "href": "",
//             "preOpenBegin": "",
//             "scheDesc": "",
//             "serviceType": "0",
//             "state": "completed",
//             "system": "unified",
//             "title": "Trading account WebSocket system upgrade",
//             "ts": "1672826038470"
//         }
//     ]
// }

func SubmitStatus() {
	ws := websocketlocal.GenWebSocket("wss://ws.okx.com:8443/ws/v5/public", 10)
	go ws.StartGather()
	ws.Submit([]byte(`{"op": "subscribe","args": [{"channel": "status"}]}`), true)
	for {
		temp := <-ws.InfoChan
		var temp_data map[string]interface{}

		json.Unmarshal(temp, &temp_data)
		if _, ok := temp_data["data"]; ok {
			SendMessage(make_formate(temp_data["data"].([]interface{})[0]), "965377515@qq.com") //15940402405@163.com
			SendMessage(make_formate(temp_data["data"].([]interface{})[0]), "15940402405@163.com")
		}
	}
}

func make_formate(info interface{}) string {
	temp_str := ""
	tool_map := map[string]string{"begin": "维护开始时间", "end": "维护结束时间", "href": "相关内容超链接", "preOpenBegin": "预开放开始的时间", "serviceType": "服务类型", "system": "系统(unified:交易账户)", "scheDesc": "改期进度说明", "ts": "推送时间"}
	tool_list := []string{"begin", "end", "href", "preOpenBegin", "serviceType", "system", "scheDesc", "ts"}
	for i := 0; i < len(tool_list); i++ {
		temp_str += "<p>"
		temp_str += tool_map[tool_list[i]]
		temp_str += ":"
		if tool_list[i] == "begin" || tool_list[i] == "end" || tool_list[i] == "ts" {
			temp_int, _ := strconv.Atoi(info.(map[string]interface{})[tool_list[i]].(string))
			temp_str += time.Unix(int64(temp_int)/1000, 0).Format("2006-01-02 15:04:05")
		} else {
			temp_str += info.(map[string]interface{})[tool_list[i]].(string)
		}
		temp_str += "</p>"
	}

	return temp_str
}

func SendMessage(test string, user string) {
	msg := gomail.NewMessage()
	//1. 设置发件人信息
	msg.SetHeader("From", "z13997171940@163.com")
	// 2. 设置收件人信息 值为  ...string 可设置多个
	msg.SetHeader("To", user) // 15940402405@163.com
	// 4. 设置邮件标题
	msg.SetHeader("Subject", "okex status update")
	// 5. 设置要发送的邮件正文
	// 第一个参数是类型，第二个参数是内容
	// 如果是 html，第一个参数则是 `text/html`
	msg.SetBody("text/html", test)
	// 传入 服务地址、端口、账号、密码 等参数，初始化 dialer实例
	dialer := gomail.NewDialer("smtp.163.com", 465, "z13997171940@163.com", "DQBAUTJFGUFJMOSP")
	// 将 msg信息 实例作为参数 传入 dialer实例，进行发送
	if err := dialer.DialAndSend(msg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("发送成功！！！")
	}
}

func main() {
	SubmitStatus()
}
