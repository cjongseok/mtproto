#!/bin/bash
cp ../types.tl.proto ../tl.proto
cp ../proxy/tl_update.proto ../proxy/update.proto
python3 -m  grpc_tools.protoc  -I .. -I ~/Programs/protoc-3.5.1/include --python_out=../py --grpc_python_out=../py ../tl.proto
python3 -m  grpc_tools.protoc -I $GOPATH/src -I ../proxy --python_out=../py --grpc_python_out=../py update.proto
rm ../tl.proto ../proxy/update.proto
