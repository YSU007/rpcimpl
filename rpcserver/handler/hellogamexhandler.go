package handler

import (
	"context"
	"log"

	"rpcimpl/rpcserver/rpcproto"
)

func HelloGameXHandler(ctx context.Context, in *rpcproto.HelloReq) (*rpcproto.HelloRsp, error) {
	log.Printf("Received: %v", in.GetName())
	return &rpcproto.HelloRsp{Message: "Hello " + in.GetName()}, nil
}
