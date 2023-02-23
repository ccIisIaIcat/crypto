package main

import (
	"fmt"
	"global"
	"log"
	"query_bar_update"
	"record_mysql_barcheck"
	"strconv"
	"strings"
	"time"
)

// 根据结尾时间和检查bar个数，生成收尾时间戳
func GenTs(time_end string, bar_count int64) (int, int) {
	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", time_end, loc)
	if err != nil {
		log.Println(err)
	}
	ts_start := (the_time.Unix() - 60*bar_count) * 1000
	ts_end := (the_time.Unix() + 1) * 1000
	return int(ts_start), int(ts_end)
}

// 从标记时间往前查找4小时
func CheckByTime(end_time string, ins_type string) {
	config := global.GetConfig("../../conf/conf.ini")
	ts_start, ts_end := GenTs(end_time, 60*4)
	check := &record_mysql_barcheck.BarCheck{}
	if ins_type == "swap" {
		check = record_mysql_barcheck.GenBarCheck(config.MysqlInfo["Local"], "crypto_swap")
	} else if ins_type == "spot" {
		check = record_mysql_barcheck.GenBarCheck(config.MysqlInfo["Local"], "crypto_spot")
	} else if ins_type == "future_quarter" {
		check = record_mysql_barcheck.GenBarCheck(config.MysqlInfo["Local"], "crypto_future_quarter")
	}
	for i := 0; i < len(check.Ins_list); i++ {
		log.Println(check.Ins_list[i], " check!")
		if check.CheckBarLose(ts_start, ts_end, check.Ins_list[i], 60*4) {
			log.Println("no error!")
		} else {
			log.Println("data missing!")
			missing_ts := check.FindMissing(ts_start, ts_end, check.Ins_list[i])
			missing_info_okex := query_bar_update.GenBarListOkex(strings.ToUpper(check.Ins_list[i]), ts_end)
			missing_info_bianace := query_bar_update.GenBarListBianace(strings.ToUpper(check.Ins_list[i]), ts_end)
			for j := 0; j < len(missing_ts); j++ {
				temp_ts := strconv.Itoa(missing_ts[j])
				if _, ok := missing_info_okex[temp_ts]; ok {
					sql := "insert into `" + strings.ToUpper(check.Ins_list[i]) + "` (Insid,Ts_open,Open_price,High_price,Low_price,Close_price,Vol,VolCcy,VolCcyQuote) values ('" + strings.ToUpper(check.Ins_list[i]) + "'," + temp_ts + "," + missing_info_okex[temp_ts][0].(string) + "," + missing_info_okex[temp_ts][1].(string) + "," + missing_info_okex[temp_ts][2].(string) + "," + missing_info_okex[temp_ts][3].(string) + "," + missing_info_okex[temp_ts][4].(string) + "," + missing_info_okex[temp_ts][5].(string) + "," + missing_info_okex[temp_ts][6].(string) + ");"
					check.InsertMissing(sql)
				} else if _, ok := missing_info_bianace[temp_ts]; ok {
					sql := "insert into `" + strings.ToUpper(check.Ins_list[i]) + "` (Insid,Ts_open,Open_price,High_price,Low_price,Close_price) values ('" + strings.ToUpper(check.Ins_list[i]) + "'," + temp_ts + "," + missing_info_bianace[temp_ts][0].(string) + "," + missing_info_bianace[temp_ts][1].(string) + "," + missing_info_bianace[temp_ts][2].(string) + "," + missing_info_bianace[temp_ts][3].(string) + ");"
					check.InsertMissing(sql)
				}
			}
			log.Println("fixed")
		}
	}
	check.Close()
}

func main() {
	// 检查的时间点和对应检查的结束时间
	check_map := map[string]string{"00:10:05": "00:00:00", "04:10:05": "04:00:00", "08:10:05": "08:00:00", "12:10:05": "12:00:00", "16:10:05": "16:00:00", "20:10:05": "20:00:00"}
	for {
		time.Sleep(time.Second)
		// "2006-01-02 15:04:05"
		temp_time := time.Now().Format("15:04:05")
		fmt.Println(temp_time)
		if _, ok := check_map[temp_time]; ok {
			temp_day := time.Now().Format("2006-01-02")
			check_time := temp_day + " " + check_map[temp_time]
			CheckByTime(check_time, "swap")
			CheckByTime(check_time, "spot")
			CheckByTime(check_time, "future_quarter")
		}
	}
}
