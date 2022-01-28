package main

import (
	"time"

	"Jottings/tiny_rpc/client"
	"Jottings/tiny_rpc/log"
	"Jottings/tiny_rpc/msg"
	"Jottings/tiny_rpc/proto"
)

const (
	network = "tcp"
	address = "localhost:8972"
)

func main() {
	log.Init(log.Zap)
	log.Info("Start..")
	defer log.Info("Stop..")
	msg.SetSerializer(msg.SerializerPB)
	var cli = client.NewClient(network, address)
	defer cli.Close()
	var req = &proto.HelloReq{
		HelloMsg: "hello",
	}
	var rsp = &proto.HelloRsp{}

	var ticker = time.NewTicker(time.Second * 2)
	for {
		select {
		case <-ticker.C:
			code, err := cli.Call(proto.Hello, req, rsp)
			if err != nil {
				return
			}
			log.Info("client call code %v rsp %v", code, rsp)
		}
	}
}
