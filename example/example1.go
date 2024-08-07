package main

import (
	"fmt"

	"github.com/skikozou/UBS/src/manager"
	"github.com/skikozou/UBS/src/server"
)

func main() {
	fmt.Println("main")

	ubs := server.New()
	ubs.HandlerFunc(func(cli *manager.Client) error {
		cli.Write("hello, world!")
		cli.Conn.Close()
		return nil
	})
	ubs.Init().Start()
}
