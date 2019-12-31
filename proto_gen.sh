#!/usr/bin/env bash

[[ ! -d internal ]] && mkdir internal
[[ ! -d internal/service ]] && mkdir internal/service

protoc --proto_path=./rpc --twirp_out=./rpc --go_out=./rpc service.proto
protoc --proto_path=./pkg/flash/model --twirp_out=./pkg/flash/model --go_out=./pkg/flash/model flash.proto