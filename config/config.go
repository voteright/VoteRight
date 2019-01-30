package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config defines the config for the VoteRight applciation
type Config struct {
	Client *ClientConfig
	Server *ServerConfig
}

// New generates a new config and loads it from the file or the environment
func New() (*Config, error) {
	cfg := &Config{}
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cfg.Client = newClientConfig(v)
	cfg.Server = newServerConfig(v)

	fmt.Printf("%+v", v.AllSettings())
	v.SetConfigType("json")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
		fmt.Println("Reading from environment")
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	fmt.Printf("Config %+v", *cfg.Server)
	return cfg, nil
}
