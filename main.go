package main

import (
	"fmt"
	"log"

	"github.com/skikozou/UBS/src/manager"
	"github.com/skikozou/UBS/src/server"
)

func main() {
	ubs := server.New()
	cfg := ubs.Init()
	cfg.SetMemoryBuffar(1024)
	cfg.SetPort("8080")
	cfg.SetisGlobal(false)
	ubs.AsyncHandlerFunc(func(resError chan<- error, cli *manager.Client) {
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
	err := make(chan error)
	go cfg.StartAsync(err)
	Err := <-err
	if Err != nil {
		log.Fatal(err)
	}
}
