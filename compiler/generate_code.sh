#!/bin/sh

#rm ../mtp/tl_schema.proto ../mtp/tl_schema.go ../proxy/procs.go
rm ../mtp/types.tl.proto ../mtp/utils.tl.go ../mtp/procs.tl.go
go run tl2go.go < tl-schema-71.json
mv types.tl.proto utils.tl.go procs.tl.go ../mtp/
protoc -I ../mtp -I ~/Programs/protoc-3.5.1/include types.tl.proto --go_out=plugins=grpc:../mtp
protoc -I $GOPATH/src -I ../proxy tl_update.proto --go_out=plugins=grpc:../proxy
go fmt ../mtp
go fmt ../proxy

#go run build_tl_scheme.go tl_schema < tl-schema-71.json  > ./tl_schema.proto
#gofmt -w ./tl_schema.go
