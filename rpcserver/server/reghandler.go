package server

import (
	"context"

	"rpcimpl/rpcserver/rpcproto"
	"rpcimpl/rpcserver/rpcserver/handler"
)

// HelloGameX implements helloworld.GreeterServer
func (s *server) HelloGameX(ctx context.Context, in *rpcproto.HelloReq) (*rpcproto.HelloRsp, error) {
	return handler.HelloGameXHandler(ctx, in)
}
