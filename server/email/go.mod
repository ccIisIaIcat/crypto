module go_email

go 1.19

require (
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	websocketlocal v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
)

replace websocketlocal => ../../websocketlocal
