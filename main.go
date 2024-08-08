package main

import (
	"github.com/skikozou/UBS/src/manager"
	"github.com/skikozou/UBS/src/server"
)

func main() {
	ubs := server.New()
	cfg := ubs.Init()
	cfg.SetMemoryBuffar(1024)
	cfg.SetPort("8585")
	cfg.SetisGlobal(true)
	ubs.ConnectEvent(func(cli *manager.Client) error {
		cli.WriteString("none")
		cli.Conn.Close()
		return nil
	})
}
