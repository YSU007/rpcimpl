package hello

import (
	"tiny_rpc/log"
	"tiny_rpc/model"
	"tiny_rpc/msg"
	"tiny_rpc/proto"
	"tiny_rpc/router"
)

type HelloProto struct {
	*proto.HelloReq
}

func (HelloProto) Serve(ctx router.ContextInterface, baseReq msg.ModeMsg, baseRsp msg.CodeMsg) {
	var account, ok = (ctx).(*model.PlayerAccount)
	if ok {
		var req = new(proto.HelloReq)
		var rsp = new(proto.HelloRsp)
		_ = msg.Unmarshal(baseReq.GetData(), req)
		var r = &HelloProto{HelloReq: req}
		var code = r.HelloHandle(account, rsp)
		var data, _ = msg.Marshal(rsp)
		baseRsp.FillIn(code, data)
	}
}

func (r *HelloProto) HelloHandle(a *model.PlayerAccount, rsp *proto.HelloRsp) (code uint32) {
	log.Info("account %v receive %v", a.AccountId, r.HelloMsg)
	rsp.ReplyMsg = "hello~"
	return
}
