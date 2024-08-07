package main

import (
	"UBS/src/manager"
	"UBS/src/server"
	"fmt"
)

func main() {
	fmt.Println("main")

	ubs := server.New()
	cfg := ubs.Init()
	cfg.SetMemoryBuffar(1024)
	cfg.SetPort("8585")
	cfg.SetisGlobal(false)
	ubs.HandlerFunc(func(cli *manager.Client) error {
		var text string
		err := cli.Read(&text, cfg.MemoryBuffer)
		if err != nil {
			return err
		}
		fmt.Println(text)

		err = cli.Write(text)
		if err != nil {
			return err
		}

		cli.Conn.Close()
		return nil
	})
	cfg.Start()
}
