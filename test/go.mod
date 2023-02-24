module demo

go 1.19

replace global => ../global

require github.com/gorilla/websocket v1.5.0

replace websocketlocal => ../websocketlocal

replace query_insid => ../query/query_insid

replace query_detail => ../query/query_detail
