package module

import "Jottings/tiny_rpc/log"

type MA struct {
}

func (r *MA) Hello(arg *HelloArg, replay *HelloReplay) error {
	log.Info("hello %v", arg.Msg)
	replay.Msg = "hello~"
	return nil
}
