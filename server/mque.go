package server

import (
	"Jottings/tiny_rpc/module"
	"Jottings/tiny_rpc/module/a"
	"Jottings/tiny_rpc/module/b"
)

var q = module.Que{
	{M: &a.MA{}, Size: module.ChanSizeDef},
	{M: &b.MB{}, Size: module.ChanSizeDef},
}
