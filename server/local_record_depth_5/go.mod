module local_record_depth_5

go 1.19

require (
	global v0.0.1
	query_depth v0.0.1
	record_mysql v0.0.1
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	websocketlocal v0.0.0-00010101000000-000000000000 // indirect
)

replace query_insid => ../../query/query_insid

replace query_depth => ../../query/query_depth

replace websocketlocal => ../../websocketlocal

replace global => ../../global

replace record_mysql => ../../record/record_mysql
