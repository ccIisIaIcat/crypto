package main

import (
	"global"
	"log"
	"record_mysql_barcheck"
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

func main() {
	config := global.GetConfig("../../conf/conf.ini")
	ts_start, ts_end := GenTs("2023-02-23 03:00:00", 60*4)
	check_spot := record_mysql_barcheck.GenBarCheck(config.MysqlInfo["Local"], "crypto_spot")
	for i := 0; i < len(check_spot.Ins_list); i++ {
		log.Println(check_spot.Ins_list[i], " check!")
		if check_spot.CheckBarLose(ts_start, ts_end, check_spot.Ins_list[i], 60*4) {
			log.Println("no error!")
		} else {
			log.Println("data missing!")
		}

	}
}
