module query_detail

go 1.19

require (
	global v0.0.0-00010101000000-000000000000
	websocketlocal v0.0.1
)

require github.com/gorilla/websocket v1.5.0 // indirect

replace websocketlocal => ../../websocketlocal

replace global => ../../global
