#!/usr/bin/env bash

# Get parent directory
DIR=$(dirname $(dirname $(readlink -fm $0)))

# Generate data models
protoc -I=$DIR/rpc/model --java_out=$DIR/rpc/jvm/src/main/java flash.proto

# Protobuf can't handle go modules
protoc -I=$DIR/rpc/model --go_out=$DIR/rpc/model flash.proto
mv $DIR/rpc/model/freggy.dev/stats/rpc/go/model/* $DIR/rpc/go/model
rm -rf $DIR/rpc/model/freggy.dev/

# Generate service libraries
protoc -I=$DIR/rpc -I=$DIR/rpc/model --twirp_out=$DIR/rpc/go/service --go_out=$DIR/rpc/go/service service.proto
protoc -I=$DIR/rpc -I=$DIR/rpc/model --java_out=$DIR/rpc/jvm/src/main/java --twirp_java_jaxrs_out=$DIR/rpc/jvm/src/main/java service.proto

