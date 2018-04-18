#!/bin/sh

rm ../types.tl.proto ../convs.tl.go ../procs.tl.go
go run tl2go.go < tl-schema-71.json
mv types.tl.proto convs.tl.go procs.tl.go ../
protoc -I .. -I ~/Programs/protoc-3.5.1/include types.tl.proto --go_out=plugins=grpc:../
protoc -I $GOPATH/src -I ../proxy tl_update.proto --go_out=plugins=grpc:../proxy
go fmt ../
go fmt ../proxy
