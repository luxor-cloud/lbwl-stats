#!/usr/bin/env bash

[[ ! -d internal ]] && mkdir internal
[[ ! -d internal/service ]] && mkdir internal/service

# Generate service libraries
protoc --proto_path=rpc --twirp_out=rpc/go/service --go_out=rpc/go/service service.proto
protoc --proto_path=rpc --java_out=rpc/jvm/src/main/java --twirp_java_jaxrs_out=rpc/jvm/src/main/java service.proto

# Generate data models
protoc -I=rpc -I=$GOPATH/src --java_out=rpc/jvm/src/main/java flash.proto
protoc -I=rpc -I=$GOPATH/src --gogo_out=rpc/go/model flash.proto