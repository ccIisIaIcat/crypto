module local_deliver

go 1.19

replace query_insid => ../../query/query_insid

replace query_tick => ../../query/query_tick

replace query_bar => ../../query/query_bar

replace query_bar_custom => ../../query/query_bar_custom

replace websocketlocal => ../../websocketlocal

replace global => ../../global

replace record_mysql => ../../record/record_mysql

replace goClient => ../../deliver/gogrpc/goClient

replace godeliver => ../../deliver/gogrpc/godeliver

replace trade_restful => ../../trade/trade_restful

replace account => ../../account

require (
	account v0.0.1
	global v0.0.1
	godeliver v0.0.1
	google.golang.org/grpc v1.53.0
	query_bar_custom v0.0.0-00010101000000-000000000000
	query_tick v0.0.0-00010101000000-000000000000
	trade_restful v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	websocketlocal v0.0.1 // indirect
)
