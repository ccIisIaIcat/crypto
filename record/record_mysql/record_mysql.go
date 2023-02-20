package record_mysql

import (
	"database/sql"
	"global"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlServer struct {
	Mysql_config global.ConfigMysql
	Stmt_map     map[string]*sql.Stmt
	db           *sql.DB
}

func GenMysqlServer(mysql_config global.ConfigMysql, database string) *MysqlServer {
	var judge global.ConfigMysql
	if judge == mysql_config {
		panic("missing config")
	}
	ms := &MysqlServer{Mysql_config: mysql_config}
	dsn := mysql_config.User + ":" + mysql_config.Password + "@tcp(" + mysql_config.Host + ":" + mysql_config.Port + ")/" + database
	var err error
	ms.db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("db format error:", err)
		panic("db format error")
	}
	err = ms.db.Ping()
	if err != nil {
		log.Println("db connecting err:", err)
		panic("db connecting err:")
	}
	log.Println("database:", database, "连接成功！")
	ms.Stmt_map = make(map[string]*sql.Stmt, 0)
	return ms
}

func (M *MysqlServer) Create_stmt(label string, sql string) {
	if _, ok := M.Stmt_map[label]; ok {
		panic("repeated key")
	}
	stmt, err := M.db.Prepare(sql)
	if err != nil {
		panic("stmt err")
	} else {
		M.Stmt_map[label] = stmt
	}
}

func (M *MysqlServer) Create_table(sql string) {
	_, err := M.db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	fmt.Println("lalala")
// 	conf := global.GetConfig("../../conf/conf.ini")
// 	fmt.Println(conf.MysqlInfo["Local"])
// 	ms := GenMysqlServer(conf.MysqlInfo["Local"], "crypto_spot")
// 	fmt.Println(ms)
// }
