package main

import (
	"fmt"
	"global"
	"query_depth"
	"record_mysql"
	"time"
)

func main() {
	conf := global.GetConfig("../../conf/conf.ini")
	// swap_insid := query_insid.GetInsByType("SWAP")
	swap_insid := []string{"BTC-USDT-SWAP", "ETH-USDT-SWAP", "LUNA-USDT-SWAP", "OKB-USDT-SWAP", "OKT-USDT-SWAP", "LTC-USDT-SWAP", "DOT-USDT-SWAP", "DOGE-USDT-SWAP", "ADA-USDT-SWAP", "XRP-USDT-SWAP"}
	ms := record_mysql.GenMysqlServer(conf.MysqlInfo["Local2"], "crypto_swap_depth_5")

	table_sql := "(id int PRIMARY KEY AUTO_INCREMENT, Insid varchar(100), Ts bigint, A1p double, A2p double, A3p double, A4p double, A5p double,A1q double, A2q double, A3q double, A4q double, A5q double, B1p double, B2p double, B3p double, B4p double, B5p double,B1q double, B2q double, B3q double, B4q double, B5q double)"
	insert_sql := "(Insid,Ts,A1p,A2p,A3p,A4p,A5p,A1q,A2q,A3q,A4q,A5q,B1p,B2p,B3p,B4p,B5p,B1q,B2q,B3q,B4q,B5q) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
	for i := 0; i < len(swap_insid); i++ {
		sql := "CREATE TABLE IF NOT EXISTS " + "`" + swap_insid[i] + "` " + table_sql + "ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;"
		ms.Create_table(sql)
	}
	// 创建对应的stmt
	for i := 0; i < len(swap_insid); i++ {
		sql := "insert into " + "`" + swap_insid[i] + "` " + insert_sql
		ms.Create_stmt(swap_insid[i], sql)
	}

	tick_info_chan := make(chan *global.Depth5Info, 100)
	qd := query_depth.Query_depth{InsId_list: swap_insid, Tick_info_chan: tick_info_chan}
	go qd.Start()
	time.Sleep(time.Second)
	for {
		temp := <-qd.Tick_info_chan
		_, err := ms.Stmt_map[temp.Insid].Exec(temp.Insid, temp.Ts_Price, temp.Ask1_price, temp.Ask2_price, temp.Ask3_price, temp.Ask4_price, temp.Ask5_price, temp.Ask1_volumn, temp.Ask2_volumn, temp.Ask3_volumn, temp.Ask4_volumn, temp.Ask5_volumn, temp.Bid1_price, temp.Bid2_price, temp.Bid3_price, temp.Bid4_price, temp.Bid5_price, temp.Bid1_volumn, temp.Bid2_volumn, temp.Bid3_volumn, temp.Bid4_volumn, temp.Bid5_volumn)
		if err != nil {
			fmt.Println(err)
		}
	}
}
