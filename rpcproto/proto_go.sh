#!/usr/bin/env bash

protoc *.proto --go_out=plugins=grpc:./
go run ../protogen/protogen.go $GOPATH
mv reghandler.go "../rpcserver/server/reghandler.go"
mv *handler.go "../rpcserver/handler/"