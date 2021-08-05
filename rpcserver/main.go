package main

import "rpcimpl/rpcserver/rpcserver/server"

const (
	netType = "tcp"
	port    = ":50051"
)

func main() {
	server.RunRPCService(netType, port)
}
