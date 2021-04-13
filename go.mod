module main

go 1.16

replace proto => ./proto

require (
	google.golang.org/grpc v1.37.0
	proto v0.0.0-00010101000000-000000000000
)
