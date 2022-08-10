#!/bin/bash

#protoc --proto_path=./pkg/ ./pkg/**/pb/*.proto --go-grpc_out=./pkg/ --go-grpc_opt=paths=source_relative

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./pkg/**/pb/*.proto