package main

import (
	"github.com/skikozou/UBS/src/manager"
	"github.com/skikozou/UBS/src/server"
)

func example1() {
	ubs := server.New()
	ubs.ConnectEvent(func(cli *manager.Client) error {
		cli.WriteString("hello, world!")
		cli.Conn.Close()
		return nil
	})
	ubs.Init().Start()
}
