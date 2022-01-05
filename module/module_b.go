package module

import "Jottings/tiny_rpc/log"

type MB struct {
}

func (r *MB) Hello() {
	log.Info("hello mb")
}
