module grcp_server

go 1.19

require (
	deliver v0.0.1
	google.golang.org/grpc v1.53.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230222225845-10f96fb3dbec // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace deliver => ./deliver
