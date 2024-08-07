package main

import (
	"UBS/src/manager"
	"UBS/src/server"
	"fmt"
	"io"
)

func main() {
	fmt.Println("main")

	ubs := server.New()
	ubs.HandlerFunc(func(cli *manager.Client) error {
		fmt.Println(cli.IP)
		_, err := io.WriteString(cli.Conn, "connected")
		return err
	})
	ubs.Init().SetPort("8585").SetisGlabal(false).Start()
}
