package server

import (
	"Jottings/tiny_rpc/handler"
	"Jottings/tiny_rpc/log"
	"Jottings/tiny_rpc/module"
	"Jottings/tiny_rpc/module/base"
	"Jottings/tiny_rpc/net"
	"Jottings/tiny_rpc/router"
)

type Server struct {
	sm *net.SessionMgr
	mm *base.ModuleMgr
}

func NewServer(network, address string) *Server {
	return &Server{
		sm: net.NewSessionMgr(network, address),
		mm: base.ModuleMgrIns(),
	}
}

func (s *Server) Serve() {
	handler.Init(router.ReflectRouterType)
	if err := s.mm.RegModule(&module.MA{}, 512); err != nil {
		log.Error("reg module err %v", err)
		return
	}
	err := s.mm.Start()
	if err != nil {
		log.Error("server module mgr start err %v", err)
		return
	}
	s.sm.Start()
}

func (s *Server) Close() {
	s.sm.Stop()
	if err := s.mm.Stop(); err != nil {
		log.Error("server module mgr stop err %v", err)
		return
	}
}
