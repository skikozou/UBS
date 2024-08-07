package main

import (
	"UBS/src/server"
	"fmt"
)

func main() {
	fmt.Println("main")

	ubs := server.New()
	//here setting
	ubs.Init().Start()
}
