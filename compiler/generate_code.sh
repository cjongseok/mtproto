#!/bin/sh

rm ../proto/tl_schema.proto ../proto/tl_schema.go
go run build_tl_scheme.go ../proto/tl_schema < tl-schema-71.json
#go run build_tl_scheme.go tl_schema < tl-schema-71.json  > ./tl_schema.proto
protoc -I ../proto -I ~/Programs/protoc-3.5.1/include tl_schema.proto --go_out=plugins=grpc:../proto
#gofmt -w ./tl_schema.go
