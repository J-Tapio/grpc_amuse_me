#!/bin/bash

mkdir -p gen/amuseme

protoc -I amuseme amuseme/amuseme.proto \
--go_out=gen/amuseme --go_opt=paths=source_relative \
--go-grpc_out=gen/amuseme --go-grpc_opt=paths=source_relative \