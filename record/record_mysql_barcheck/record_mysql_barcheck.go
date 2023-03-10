package record_mysql_barcheck

import (
	"database/sql"
	"global"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type BarCheck struct {
	Mysql_config global.ConfigMysql
	Stmt_map     map[string]*sql.Stmt
	db           *sql.DB
	database     string
	Ins_list     []string
}

func GenBarCheck(mysql_config global.ConfigMysql, database string) *BarCheck {
	bc := &BarCheck{}
	bc.database = database
	bc.Mysql_config = mysql_config
	dsn := mysql_config.User + ":" + mysql_config.Password + "@tcp(" + mysql_config.Host + ":" + mysql_config.Port + ")/" + database
	var err error
	bc.db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("db format error:", err)
		panic("db format error")
	}
	err = bc.db.Ping()
	if err != nil {
		log.Println("db connecting err:", err)
		panic("db connecting err:")
	}
	log.Println("database:", database, "连接成功！")
	bc.Ins_list = bc.GetInsid()

	return bc
}

func (B *BarCheck) GetInsid() []string {
	rows, err := B.db.Query("SHOW TABLES;")
	answer := make([]string, 0)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var temp string
		rows.Scan(&temp)
		answer = append(answer, temp)
	}
	return answer

}

// 给定首尾时间戳和insid，检查是否丢包
func (B *BarCheck) CheckBarLose(ts_start int, ts_end int, insid string, check_sum int) bool {
	t1 := strconv.Itoa(ts_start)
	t2 := strconv.Itoa(ts_end)
	row := B.db.QueryRow("select count(*) from `" + insid + "` where Ts_open > " + t1 + " and Ts_open <" + t2 + ";")
	var temp int
	row.Scan(&temp)
	if temp == check_sum {
		return true
	} else {
		return false
	}

}

// 给定首尾时间戳和insid,查找丢失的包
func (B *BarCheck) FindMissing(ts_start int, ts_end int, insid string) []int {
	t1 := strconv.Itoa(ts_start)
	t2 := strconv.Itoa(ts_end)
	tool_map := make(map[int]bool, 0)
	temp_ts := ts_start
	for temp_ts < ts_end {
		tool_map[temp_ts] = false
		temp_ts += 60 * 1000
	}
	rows, _ := B.db.Query("select Ts_open from `" + insid + "` where Ts_open > " + t1 + " and Ts_open <" + t2 + ";")
	var temp int
	for rows.Next() {
		rows.Scan(&temp)
		tool_map[temp] = true
	}
	answer := make([]int, 0)
	for k, v := range tool_map {
		if !v {
			answer = append(answer, k)
		}
	}
	return answer
}

// 在找到缺失的条和可以填充的条后，进行填充,注明是从okex获取的还是从币安获取的，应为币安不填充交易量
func (B *BarCheck) InsertMissing(sql string) {
	if sql != "" {
		_, err := B.db.Exec(sql)
		if err != nil {
			log.Println(err)
		}
	}

}

func (B *BarCheck) Close() {
	B.db.Close()
}

// func main() {
// 	user_conf := global.GetConfig("../../conf/conf.ini")
// 	bc := GenBarCheck(user_conf.MysqlInfo["Local"], "crypto_swap")
// 	/////
// 	loc, _ := time.LoadLocation("Local")
// 	stringTime := "2023-02-23 10:40:00"
// 	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	ts_start := (the_time.Unix() - 3600*4) * 1000
// 	ts_end := (the_time.Unix() + 1) * 1000
// 	if bc.CheckBarLose(int(ts_start), int(ts_end), "cvc-usdt-swap", 240) {
// 		log.Println("check no error")
// 	} else {
// 		log.Println("error")
// 	}
// }
