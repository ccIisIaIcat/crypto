module local_deliver

go 1.19

replace query_insid => ../../query/query_insid

replace query_tick => ../../query/query_tick

replace websocketlocal => ../../websocketlocal

replace global => ../../global

replace record_mysql => ../../record/record_mysql

replace goClient => ../../deliver/gogrpc/goClient

replace godeliver => ../../deliver/gogrpc/godeliver

require (
	global v0.0.1
	goClient v0.0.0-00010101000000-000000000000
	query_tick v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	godeliver v0.0.1 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	websocketlocal v0.0.0-00010101000000-000000000000 // indirect
)
