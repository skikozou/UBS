package manager

import "net"

type Client struct {
	Conn *net.TCPConn
	IP   string
	Port string
}
