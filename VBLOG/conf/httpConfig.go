package conf

import "fmt"

type Http struct {
	Host string `json:"host" toml:"host" env:"HTTP_HOST"`
	Port int    `json:"port" toml:"port" env:"HTTP_PORT"`
}

func newDefaultHttpConfig() *Http {
	return &Http{
		Host: "0.0.0.0",
		Port: 11220,
	}
}

func (h *Http) Address() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}
