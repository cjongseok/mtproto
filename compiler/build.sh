#!/bin/sh

rm ../core/types.tl.proto ../core/convs.tl.go ../core/procs.tl.go
go run tl2go.go < tl-schema-71.json
mv types.tl.proto convs.tl.go procs.tl.go ../core/
protoc -I ../core -I ~/Programs/protoc-3.5.1/include types.tl.proto --go_out=plugins=grpc:../core
protoc -I $GOPATH/src -I ../proxy tl_update.proto --go_out=plugins=grpc:../proxy
go fmt ../core
go fmt ../proxy
