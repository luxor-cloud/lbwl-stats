#!/usr/bin/env bash

# Get parent directory
DIR=$(dirname $(dirname $(readlink -fm $0)))

# Generate service libraries
protoc -I=$DIR/rpc --twirp_out=$DIR/rpc/go/service --go_out=$DIR/rpc/go/service service.proto
protoc -I=$DIR/rpc --java_out=$DIR/rpc/jvm/src/main/java --twirp_java_jaxrs_out=$DIR/rpc/jvm/src/main/java service.proto

# Generate data models
protoc -I=$DIR/rpc -I=$GOPATH/src --java_out=$DIR/rpc/jvm/src/main/java flash.proto
protoc -I=$DIR/rpc -I=$GOPATH/src --go_out=$DIR/rpc/go/model flash.proto
