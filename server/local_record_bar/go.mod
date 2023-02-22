module local_record_bar

go 1.19

require query_insid v0.0.1

require (
	global v0.0.1
	query_bar v0.0.1
	record_mysql v0.0.1
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	websocketlocal v0.0.1 // indirect
)

replace query_insid => ../../query/query_insid

replace query_bar => ../../query/query_bar

replace websocketlocal => ../../websocketlocal

replace global => ../../global

replace record_mysql => ../../record/record_mysql
