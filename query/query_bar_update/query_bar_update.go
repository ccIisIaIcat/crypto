package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 指定insid，和时间节点，返回时间节点二十四小时内的bar数据(从okex获取)
func GenBarListOkex(insid string, ts int) map[string]([]interface{}) {
	answer := make(map[string]([]interface{}), 0)
	time_after := strconv.Itoa((ts) + 1)
	data, err := http.Get("https://www.okx.com/api/v5/market/candles?instId=" + insid + "&after=" + time_after + "&limit=240")
	if err != nil {
		fmt.Println(err)
	}
	temp_json, _ := io.ReadAll(data.Body)
	var temp map[string](interface{})
	json.Unmarshal(temp_json, &temp)
	// fmt.Println(temp["data"].([]interface{}))
	info_list := temp["data"].([]interface{})
	for i := 0; i < len(info_list); i++ {
		tt, _ := strconv.Atoi(info_list[i].([]interface{})[0].(string))
		tm := time.Unix(int64(tt/1000), 0)
		fmt.Println(tm)
		answer[info_list[i].([]interface{})[0].(string)] = info_list[i].([]interface{})[1:]
	}
	return answer
}

// 指定insid，和时间节点，返回时间节点二十四小时内的bar数据(从bianace获取)[ps:局限于现货]
func GenBarListBianace(insid string, ts int) map[string]([]interface{}) {
	answer := make(map[string]([]interface{}), 10)
	insid = strings.Join(strings.Split(insid, "-"), "")
	time_after := strconv.Itoa((ts) + 1)
	data, err := http.Get("https://api4.binance.com/api/v3/klines?symbol=" + insid + "&interval=1m&limit=240&endTime=" + time_after)
	if err != nil {
		fmt.Println(err)
	}
	temp_json, _ := io.ReadAll(data.Body)
	var temp ([]interface{})
	json.Unmarshal(temp_json, &temp)
	for i := 0; i < len(temp); i++ {
		temp_list := temp[i].([]interface{})
		tm := time.Unix(int64(temp_list[0].(float64)/1000), 0)
		fmt.Println(tm)
		answer[strconv.Itoa(int(temp_list[0].(float64)))] = append(temp_list[1:6], temp_list[7:9]...)
	}
	return answer

}

func main() {
	loc, _ := time.LoadLocation("Local")
	stringTime := "2023-02-22 17:00:00"
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(the_time.Unix() * 1000)
	lala := GenBarListOkex("ETH-USDT", int(the_time.Unix()*1000))

	fmt.Println(len(lala))

}
