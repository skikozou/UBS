package main

import (
	"github.com/skikozou/UBS/src/manager"
	"github.com/skikozou/UBS/src/server"
)

func example3() {
	ubs := server.New()
	cfg := ubs.Init()
	cfg.Port = "8080"
	cfg.SetisGlobal(true)
	cfg.MemoryBuffer = 1024
	ubs.HandlerFunc(func(cli *manager.Client) error {
		cli.WriteString("hello, world!")
		cli.Conn.Close()
		return nil
	})
	cfg.Start()
}
