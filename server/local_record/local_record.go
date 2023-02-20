package main

import (
	"global"
	"query_detail"
	"query_insid"
	"record_mysql"
	"time"
)

func GatherSwap() {
	// 查询swap合约代码
	swap_insid := query_insid.GetInsByType("SWAP")

	conf := global.GetConfig("../../conf/conf.ini")
	ms := record_mysql.GenMysqlServer(conf.MysqlInfo["Local"], "crypto_swap")
	// 检查对应表是否创建
	for i := 0; i < len(swap_insid); i++ {
		sql := "CREATE TABLE IF NOT EXISTS " + "`" + swap_insid[i] + "`" + "(id int PRIMARY KEY AUTO_INCREMENT, Insid varchar(100), Ts_price bigint, Ask1_price double, Bid1_price double, Oi bigint, OiCcy double, Ts_OpenInterest bigint, FundingRate double, NextFundingRate double, Ts_FundingRate bigint, TS_NextFundingRate bigint)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;"
		ms.Create_table(sql)
	}
	// 创建对应的stmt
	for i := 0; i < len(swap_insid); i++ {
		sql := "insert into " + "`" + swap_insid[i] + "`" + " (Insid,Ts_price,Ask1_price,Bid1_price,Oi,OiCcy,Ts_OpenInterest,FundingRate,NextFundingRate,Ts_FundingRate,TS_NextFundingRate) values(?,?,?,?,?,?,?,?,?,?,?);"
		ms.Create_stmt(swap_insid[i], sql)
	}

	query_obj := query_detail.QueryDetail{InsId_list: swap_insid}
	go query_obj.Start()
	time.Sleep(time.Second)
	for {
		temp := <-query_obj.Tick_info_chan
		ms.Stmt_map[temp.Insid].Exec(temp.Insid, temp.Ts_Price, temp.Ask1_price, temp.Bid1_price, temp.Oi, temp.OiCcy, temp.Ts_OpenInterest, temp.FundingRate, temp.NextFundingRate, temp.Ts_FundingRate, temp.TS_NextFundingRate)
	}
}

func main() {
	go GatherSwap()
	global.NeverStop()
}
