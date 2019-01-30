package config

import (
	"github.com/spf13/viper"
)

// ServerConfig defines the config for the vote counting server, including the addresses of other servers
type ServerConfig struct {
	Port string
}

func newServerConfig(v *viper.Viper) *ServerConfig {
	c := &ServerConfig{
		Port: "8080",
	}
	v.SetDefault("server.port", "8080")
	return c
}
