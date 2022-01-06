package a

import (
	"Jottings/tiny_rpc/log"
	"Jottings/tiny_rpc/module"
	"Jottings/tiny_rpc/proto"
)

type MA struct {
}

func (r *MA) Load() error {
	panic("implement me")
}

func (r *MA) Save() error {
	panic("implement me")
}

func (r *MA) Hello(arg *proto.HelloArg, replay *proto.HelloReplay) error {
	log.Info("hello %v", arg.Msg)
	replay.Msg = "hello~"
	module.NotifyWork("MB", "Hello", &proto.HelloBArg{Msg: "mb"})
	return nil
}
