package main

import (
	"fmt"

	"github.com/skikozou/UBS/src/manager"
	"github.com/skikozou/UBS/src/server"
)

func example3() {
	ubs := server.New()
	cfg := ubs.Init()
	cfg.Port = "8080"
	cfg.SetisGlobal(true)
	cfg.MemoryBuffer = 1024
	ubs.ConnectEvent(func(cli *manager.Client) error {
		var text string
		err := cli.ReadString(&text, cfg.MemoryBuffer)
		if err != nil {
			return err
		}
		fmt.Println(text)

		err = cli.WriteBytes([]byte(text))
		if err != nil {
			return err
		}

		cli.Conn.Close()
		return nil
	})
	cfg.Start()
}
