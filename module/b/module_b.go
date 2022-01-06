package b

import (
	"Jottings/tiny_rpc/log"
	"Jottings/tiny_rpc/proto"
)

type MB struct {
}

func (r *MB) Load() error {
	panic("implement me")
}

func (r *MB) Save() error {
	panic("implement me")
}

func (r *MB) Hello(arg *proto.HelloBArg) {
	log.Info("hello %s", arg.Msg)
}
