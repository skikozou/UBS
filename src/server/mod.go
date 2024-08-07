package server

import (
	"UBS/src/manager"
	"UBS/src/server"
	"net"
	"strings"
)

type UBS struct {
	Engine *Engine
}

func New() *UBS {
	return &UBS{
		Engine: &Engine{
			Config: &EngineConfig{},
		},
	}
}

//listen function

type EngineConfig struct {
	Port         string
	isPortReady  bool
	isGlobal     bool
	MemoryBuffer int
	Start        func() error
	Err          error
}

func (c *EngineConfig) SetPort(port string) *EngineConfig {
	c.Port = ":" + port
	return c
}

//Init

func (u *UBS) Init() *EngineConfig {
	cfg := &EngineConfig{
		Port:         ":8080",
		isPortReady:  true,
		isGlobal:     true,
		MemoryBuffer: 1024,
	}
	cfg.Start = func() error {
		return u.Engine.Run(cfg.Port, cfg.MemoryBuffer)
	}

	return cfg
}

//onRequest

func (e *Engine) onRequest(lisn *net.TCPListener) error {
	for {
		conn, err := lisn.AcceptTCP()
		if err != nil {
			return err
		}
		addr := strings.Split(conn.RemoteAddr().String(), ":")
		cli := &manager.Client{
			Conn: conn,
			IP:   addr[0],
			Port: ":" + addr[1],
		}
		return server.Request(cli)
	}
}
