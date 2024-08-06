package server

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
	Port        string
	isPortReady bool

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
		MemoryBuffer: 1024,
	}
	cfg.Start = func() error {
		return u.Engine.Run(cfg.Port, cfg.MemoryBuffer)
	}

	return cfg
}
