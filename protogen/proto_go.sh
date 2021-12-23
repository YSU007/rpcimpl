#!/bin/sh

PROTO_PATH=./
OUT_PATH=../proto

protoc --proto_path="$PROTO_PATH" --go_out=$OUT_PATH ./*.proto