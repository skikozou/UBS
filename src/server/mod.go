package server

import (
	"fmt"

	"github.com/skikozou/UBS/src/manager"
)

type UBS struct {
	Engine *Engine
}

func New() *UBS {
	return &UBS{
		Engine: &Engine{
			Config:   &EngineConfig{},
			ExitFlag: false,
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
	Handler      func(cli *manager.Client) error
}

func (c *EngineConfig) SetPort(port string) *EngineConfig {
	c.Port = ":" + port
	return c
}

func (c *EngineConfig) SetisGlobal(isglobal bool) *EngineConfig {
	c.isGlobal = isglobal
	return c
}

func (c *EngineConfig) SetMemoryBuffar(buf int) *EngineConfig {
	c.MemoryBuffer = buf
	return c
}

func (c *EngineConfig) DebugCfg() *EngineConfig {
	fmt.Println(c)
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

//Handler Func

func (u *UBS) ConnectEvent(req Request) *UBS {
	u.Engine.Config.Handler = req
	return u
}
