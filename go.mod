module main

go 1.16

replace local/pokemon => ./pokemon

require (
	google.golang.org/grpc v1.37.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	local/pokemon v0.0.0-00010101000000-000000000000
)
