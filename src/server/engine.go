package server

import (
	"net"
	"strings"

	"github.com/skikozou/UBS/src/manager"
)

type Engine struct {
	Config   *EngineConfig
	ExitFlag bool
}

type Request func(cli *manager.Client) error
type AsyncRequest func(resError chan<- error, cli *manager.Client)

func (e *Engine) Run(port string, mbuf int) error {
	host := ""
	if e.Config.isGlobal {
		host = "0.0.0.0"
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+port)
	if err != nil {
		return err
	}
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	for {
		if e.ExitFlag {
			return nil
		}
		conn, err := ln.AcceptTCP()
		if err != nil {
			return err
		}
		addr := strings.Split(conn.RemoteAddr().String(), ":")
		cli := &manager.Client{
			Conn: conn,
			IP:   addr[0],
			Port: ":" + addr[1],
		}
		err = e.Config.Handler(cli)
		if err != nil {
			return err
		}
	}
}

func (e *Engine) Exit() {
	e.ExitFlag = true
}
