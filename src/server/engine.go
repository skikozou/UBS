package server

import (
	"UBS/src/manager"
	"net"
)

type Engine struct {
	Config *EngineConfig
}

func (e *Engine) Run(port string, mbuf int) error {
	host := ""
	if e.Config.isGlobal {
		host = "0.0.0.0"
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+e.Config.Port)
	if err != nil {
		return err
	}
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	return e.onRequest(ln)
}

type Request func(cli *manager.Client)
