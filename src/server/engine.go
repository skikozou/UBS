package server

import (
	"net"
)

type Engine struct {
	Config *EngineConfig
}

func (e *Engine) Run(port string, mbuf int) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", e.Config.Port)
	if err != nil {
		return err
	}
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	conn, err := ln.AcceptTCP()
	//受け付けた後の関数
}
