package model

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type CliConfig struct {
	baseUrl string
	Debug   bool
	Profile string
}

func (c CliConfig) BaseUrl() string {
	return c.baseUrl
}

func NewCliConfig(flags pflag.FlagSet, viper viper.Viper) CliConfig {
	profile, _ := flags.GetString("profile")
	debug, _ := flags.GetBool("debug")
	config := CliConfig{
		Debug:   debug,
		Profile: profile,
	}

	cfg := viper.Sub(profile)
	if cfg != nil {
		config.baseUrl = cfg.GetString("base-url")
	}

	return config
}
