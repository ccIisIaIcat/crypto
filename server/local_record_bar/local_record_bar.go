package main

import (
	"global"
	"log"
	"query_bar"
	"query_insid"
	"record_mysql"
	"time"
)

// type BarInfo struct {
// 	// bar信息
// 	Insid       string
// 	Ts_open     int
// 	Open_price  float64
// 	High_price  float64
// 	Low_price   float64
// 	Close_price float64
// 	Vol         float64 // 交易量，以张为单位
// 	VolCcy      float64 // 交易量，以币为单位
// 	VolCcyQuote float64 // 交易量，以计价货币为单位
// 	// 如果是SWAP，还保存资金费率信息
// 	FundingRate        float64 // 当前资金费率
// 	NextFundingRate    float64 // 下一期预测资金费率
// 	Ts_FundingRate     int     // 资金费率最后更新时间
// 	TS_NextFundingRate int
// }

func GatherSwap() {
	// 查询swap合约代码
	swap_insid := query_insid.GetInsByType("SWAP")

	conf := global.GetConfig("../../conf/conf.ini")
	ms := record_mysql.GenMysqlServer(conf.MysqlInfo["Local"], "crypto_swap")
	// 检查对应表是否创建
	for i := 0; i < len(swap_insid); i++ {
		sql := "CREATE TABLE IF NOT EXISTS " + "`" + swap_insid[i] + "`" + "(id int PRIMARY KEY AUTO_INCREMENT, Insid varchar(100), Ts_open bigint, Open_price double, High_price double, Low_price double, Close_price double, Vol double, VolCcy double, VolCcyQuote double, FundingRate double, NextFundingRate double, Ts_FundingRate bigint, TS_NextFundingRate bigint)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;"
		ms.Create_table(sql)
	}
	// 创建对应的stmt
	for i := 0; i < len(swap_insid); i++ {
		sql := "insert into " + "`" + swap_insid[i] + "`" + " (Insid,Ts_open,Open_price,High_price,Low_price,Close_price,Vol,VolCcy,VolCcyQuote,FundingRate,NextFundingRate,Ts_FundingRate,TS_NextFundingRate) values(?,?,?,?,?,?,?,?,?,?,?,?,?);"
		ms.Create_stmt(swap_insid[i], sql)
	}

	query_obj := query_bar.QueryBar{InsId_list: swap_insid, InsType: "SWAP"}
	go query_obj.Start()
	time.Sleep(time.Second)
	for {
		temp := <-query_obj.Bar_info_chan
		_, err := ms.Stmt_map[temp.Insid].Exec(temp.Insid, temp.Ts_open, temp.Open_price, temp.High_price, temp.Low_price, temp.Close_price, temp.Vol, temp.VolCcy, temp.VolCcyQuote, temp.FundingRate, temp.NextFundingRate, temp.Ts_FundingRate, temp.TS_NextFundingRate)
		if err != nil {
			log.Println(err)
		}
	}
}

func GatherSpot() {
	spot_insid := query_insid.GetInsByType("SPOT")
	conf := global.GetConfig("../../conf/conf.ini")
	ms := record_mysql.GenMysqlServer(conf.MysqlInfo["Local"], "crypto_spot")
	// 检查对应表是否创建
	for i := 0; i < len(spot_insid); i++ {
		sql := "CREATE TABLE IF NOT EXISTS " + "`" + spot_insid[i] + "`" + "(id int PRIMARY KEY AUTO_INCREMENT, Insid varchar(100), Ts_open bigint, Open_price double, High_price double, Low_price double, Close_price double, Vol double, VolCcy double, VolCcyQuote double)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;"
		ms.Create_table(sql)
	}
	// 创建对应的stmt
	for i := 0; i < len(spot_insid); i++ {
		sql := "insert into " + "`" + spot_insid[i] + "`" + " (Insid,Ts_open,Open_price,High_price,Low_price,Close_price,Vol,VolCcy,VolCcyQuote) values(?,?,?,?,?,?,?,?,?);"
		ms.Create_stmt(spot_insid[i], sql)
	}
	query_obj := query_bar.QueryBar{InsId_list: spot_insid, InsType: "SPOT"}
	go query_obj.Start()
	time.Sleep(time.Second)
	for {
		temp := <-query_obj.Bar_info_chan
		ms.Stmt_map[temp.Insid].Exec(temp.Insid, temp.Ts_open, temp.Open_price, temp.High_price, temp.Low_price, temp.Close_price, temp.Vol, temp.VolCcy, temp.VolCcyQuote)
	}
}

func GatherFutureQuarter() {
	future_insid := query_insid.GetInsByTypeAndAlias("FUTURES", "quarter")
	conf := global.GetConfig("../../conf/conf.ini")
	ms := record_mysql.GenMysqlServer(conf.MysqlInfo["Local"], "crypto_future_quarter")
	// 检查对应表是否创建
	for i := 0; i < len(future_insid); i++ {
		sql := "CREATE TABLE IF NOT EXISTS " + "`" + future_insid[i] + "`" + "(id int PRIMARY KEY AUTO_INCREMENT, Insid varchar(100), Ts_open bigint, Open_price double, High_price double, Low_price double, Close_price double, Vol double, VolCcy double, VolCcyQuote double)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;"
		ms.Create_table(sql)
	}
	// 创建对应的stmt
	for i := 0; i < len(future_insid); i++ {
		sql := "insert into " + "`" + future_insid[i] + "`" + " (Insid,Ts_open,Open_price,High_price,Low_price,Close_price,Vol,VolCcy,VolCcyQuote) values(?,?,?,?,?,?,?,?,?);"
		ms.Create_stmt(future_insid[i], sql)
	}
	query_obj := query_bar.QueryBar{InsId_list: future_insid, InsType: "SPOT"}
	go query_obj.Start()
	time.Sleep(time.Second)
	for {
		temp := <-query_obj.Bar_info_chan
		_, err := ms.Stmt_map[temp.Insid].Exec(temp.Insid, temp.Ts_open, temp.Open_price, temp.High_price, temp.Low_price, temp.Close_price, temp.Vol, temp.VolCcy, temp.VolCcyQuote)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	go GatherSwap()
	go GatherSpot()
	go GatherFutureQuarter()

	global.NeverStop()
}
