package config

import (
	"github.com/spf13/viper"
)

// ClientConfig defines the config for the voting service, which handles the casting of user votes
type ClientConfig struct {
}

func newClientConfig(v *viper.Viper) *ClientConfig {
	return nil
}
