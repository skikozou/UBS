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
	cfg.SetPort("8585")
	cfg.SetisGlobal(true)
	ubs.HandlerAsyncFunc(func(resError chan<- error, cli *manager.Client) {
		var text string
		err := cli.ReadString(&text, cfg.MemoryBuffer)
		if err != nil {
			resError <- err
			return
		}
		fmt.Println(text)

		err = cli.WriteString(text)
		cli.Conn.Close()
		resError <- err
	})
	err := make(chan error)
	go cfg.StartAsync(err)
	fmt.Println("Async module")
	Err := <-err
	if Err != nil {
		log.Fatal(err)
	}
}
