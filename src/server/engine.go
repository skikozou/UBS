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
type AsyncRequest func(resError chan<- error, cli *manager.Client) error

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

func (e *Engine) RunAsync(resError chan<- error, port string, mbuf int) {
	nError := make(chan error)
	host := ""
	if e.Config.isGlobal {
		host = "0.0.0.0"
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+port)
	if err != nil {
		resError <- err
		return
	}
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		resError <- err
		return
	}
	for {
		if e.ExitFlag {
			resError <- nil
			return
		}
		conn, err := ln.AcceptTCP()
		if err != nil {
			resError <- err
			return
		}
		addr := strings.Split(conn.RemoteAddr().String(), ":")
		cli := &manager.Client{
			Conn: conn,
			IP:   addr[0],
			Port: ":" + addr[1],
		}
		e.Config.AsyncHandler(nError, cli)
		err = <-nError
		if err != nil {
			resError <- err
			return
		}
	}
}

func (e *Engine) Exit() {
	e.ExitFlag = true
}
