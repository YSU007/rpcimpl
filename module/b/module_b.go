package b

import (
	"tiny_rpc/log"
	"tiny_rpc/proto"
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
