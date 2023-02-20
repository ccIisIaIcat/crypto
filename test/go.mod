module demo

go 1.19

replace global => ../global

require query_detail v0.0.1

require query_insid v0.0.1

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	global v0.0.0-00010101000000-000000000000 // indirect
	websocketlocal v0.0.1 // indirect
)

replace websocketlocal => ../websocketlocal

replace query_insid => ../query/query_insid

replace query_detail => ../query/query_detail
