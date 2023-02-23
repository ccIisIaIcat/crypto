package record_mysql_barcheck

import (
	"database/sql"
	"fmt"
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
	fmt.Println(temp)
	if temp == check_sum {
		return true
	} else {
		return false
	}

}

// 在找到缺失的条和可以填充的条后，进行填充,注明是从okex获取的还是从币安获取的，应为币安不填充交易量
func (B *BarCheck) InsertMissing(barinfo global.BarInfo, exchange string) {
	sql := ""
	if exchange == "okex" {
		strconv.Itoa(barinfo.Ts_open)
		strconv.FormatFloat(barinfo.Open_price, 'f', 1, 64)
		strconv.FormatFloat(barinfo.High_price, 'f', 1, 64)
		strconv.FormatFloat(barinfo.Low_price, 'f', 1, 64)
		strconv.FormatFloat(barinfo.Close_price, 'f', 1, 64)
		strconv.FormatFloat(barinfo.Vol, 'f', 1, 64)
		strconv.FormatFloat(barinfo.VolCcy, 'f', 1, 64)
		strconv.FormatFloat(barinfo.VolCcyQuote, 'f', 1, 64)
		sql = "insert into `" + barinfo.Insid + "`(Insid,Ts_open,Open_price,High_price,Low_price,Close_price,Vol,VolCcy,VolCcyQuote) values(" + barinfo.Insid + strconv.Itoa(barinfo.Ts_open) + strconv.FormatFloat(barinfo.Open_price, 'f', 1, 64) + strconv.FormatFloat(barinfo.High_price, 'f', 1, 64) + strconv.FormatFloat(barinfo.Low_price, 'f', 1, 64) + strconv.FormatFloat(barinfo.Close_price, 'f', 1, 64) + strconv.FormatFloat(barinfo.Vol, 'f', 1, 64) + strconv.FormatFloat(barinfo.VolCcy, 'f', 1, 64) + strconv.FormatFloat(barinfo.VolCcyQuote, 'f', 1, 64) + ");"
	} else if exchange == "bianace" {
		strconv.Itoa(barinfo.Ts_open)
		strconv.FormatFloat(barinfo.Open_price, 'f', 1, 64)
		strconv.FormatFloat(barinfo.High_price, 'f', 1, 64)
		strconv.FormatFloat(barinfo.Low_price, 'f', 1, 64)
		strconv.FormatFloat(barinfo.Close_price, 'f', 1, 64)
		sql = "insert into `" + barinfo.Insid + "`(Insid,Ts_open,Open_price,High_price,Low_price,Close_price) values(" + barinfo.Insid + strconv.Itoa(barinfo.Ts_open) + strconv.FormatFloat(barinfo.Open_price, 'f', 1, 64) + strconv.FormatFloat(barinfo.High_price, 'f', 1, 64) + strconv.FormatFloat(barinfo.Low_price, 'f', 1, 64) + strconv.FormatFloat(barinfo.Close_price, 'f', 1, 64) + ");"
	}
	if sql != "" {
		_, err := B.db.Exec(sql)
		if err != nil {
			log.Println(err)
		}
	}

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
