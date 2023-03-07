module account

go 1.19

require websocketlocal v0.0.1

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)

require global v0.0.1

replace websocketlocal => ../websocketlocal

replace global => ../global
