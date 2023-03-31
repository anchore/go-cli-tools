package config

import "github.com/spf13/pflag"

type Config struct {
	AppName    string `json:"-" yaml:"-" mapstructure:"-"`
	ConfigFile string `json:"config,omitempty" yaml:"config,omitempty" mapstructure:"-"`
}

func NewConfig(appName string) *Config {
	return &Config{
		AppName: appName,
	}
}

func (r *Config) AddFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&r.ConfigFile, "config", "c", r.ConfigFile, "configuration file")
}
