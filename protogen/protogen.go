package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/emicklei/proto"
)

var (
	gopath     = ""
	proto_path = "/src/rpcimpl/rpcproto/msg.proto"
	out_path   = "/src/rpcimpl/rpcproto/"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("gopath none")
		return
	}
	gopath = os.Args[1]
	proto_path = gopath + proto_path
	out_path = gopath + out_path
	fmt.Println("GOPATH:", gopath)

	reader, _ := os.Open(proto_path)
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, _ := parser.Parse()

	proto.Walk(definition,
		proto.WithService(handleService))
}

func handleService(s *proto.Service) {
	const all = -1
	regStr := HANDLE_REG_HEAD_TMPL
	for _, e := range s.Elements {
		rpc, ok := e.(*proto.RPC)
		if !ok {
			continue
		}
		handleRegStr := HANDLE_REG_TMPL
		handleRegStr = strings.Replace(handleRegStr, HANDLE_NAME, rpc.Name, all)
		handleRegStr = strings.Replace(handleRegStr, HANDLE_REQ_NAME, rpc.RequestType, all)
		handleRegStr = strings.Replace(handleRegStr, HANDLE_RSP_NAME, rpc.ReturnsType, all)
		regStr += handleRegStr

		handleStr := HANDLE_TMPL
		handleStr = strings.Replace(handleStr, HANDLE_NAME, rpc.Name, all)
		handleStr = strings.Replace(handleStr, HANDLE_REQ_NAME, rpc.RequestType, all)
		handleStr = strings.Replace(handleStr, HANDLE_RSP_NAME, rpc.ReturnsType, all)
		handleFile := strings.ToLower(rpc.Name) + "handler.go"
		f, _ := os.Create(out_path + handleFile)
		f.Write([]byte(handleStr))
		f.Close()
		fmt.Println("RPC Service", s.Name, handleFile, "gen")
	}
	regFile := "reghandler.go"
	f, _ := os.Create(out_path + regFile)
	f.Write([]byte(regStr))
	f.Close()
	fmt.Println("RPC Service", s.Name, regFile, "gen")
}

const (
	HANDLE_NAME     = `{HANDLE_NAME}`
	HANDLE_REQ_NAME = `{HANDLE_REQ_NAME}`
	HANDLE_RSP_NAME = `{HANDLE_RSP_NAME}`
	HANDLE_TMPL     = `package handler

import (
	"context"
	"rpcimpl/rpcserver/rpcproto"
)

func {HANDLE_NAME}Handler(ctx context.Context, in *rpcproto.{HANDLE_REQ_NAME}, out *rpcproto.{HANDLE_RSP_NAME}) error {
    return nil
}
`
	HANDLE_REG_HEAD_TMPL = `package server

import (
	"context"

	"rpcimpl/rpcserver/rpcproto"
	"rpcimpl/rpcserver/rpcserver/handler"
)
`
	HANDLE_REG_TMPL = `
func (s *server) {HANDLE_NAME}(ctx context.Context, in *rpcproto.{HANDLE_REQ_NAME}) (*rpcproto.{HANDLE_RSP_NAME}, error) {
	var out = &rpcproto.{HANDLE_RSP_NAME}{}
	return out, handler.{HANDLE_NAME}Handler(ctx, in, out)
}
`
)
