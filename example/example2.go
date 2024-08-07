package main

import (
	"github.com/skikozou/UBS/src/manager"
	"github.com/skikozou/UBS/src/server"
)

func example2() {
	ubs := server.New()
	ubs.HandlerFunc(func(cli *manager.Client) error {
		cli.Write("hello, world!")
		cli.Conn.Close()
		return nil
	})
	ubs.Init().SetPort("8668").SetisGlobal(true).SetMemoryBuffar(1024).Start()
}
