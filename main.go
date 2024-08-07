package main

import (
	"fmt"

	"github.com/skikozou/UBS/src/manager"
	"github.com/skikozou/UBS/src/server"
)

func main() {
	ubs := server.New()
	cfg := ubs.Init()
	cfg.SetMemoryBuffar(1024)
	cfg.SetPort("8585")
	cfg.SetisGlobal(false)
	ubs.HandlerFunc(func(cli *manager.Client) error {
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
