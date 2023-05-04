package log

import (
	"github.com/spf13/pflag"

	"github.com/anchore/fangs/config"
)

type Level string

const (
	DisabledLevel Level = ""
	ErrorLevel    Level = "error"
	WarnLevel     Level = "warn"
	InfoLevel     Level = "info"
	DebugLevel    Level = "debug"
	TraceLevel    Level = "trace"
)

type Config struct {
	Level     Level  `json:"level" yaml:"level" mapstructure:"level"`
	Verbosity int    `json:"verbosity" yaml:"verbosity" mapstructure:"verbosity"`
	File      string `json:"file" yaml:"file" mapstructure:"file"`
	Quiet     bool   `json:"quiet" yaml:"quiet" mapstructure:"quiet"`
}

func NewConfig() *Config {
	return &Config{
		Level:     WarnLevel,
		Verbosity: 1,
		File:      "",
		Quiet:     false,
	}
}

func (c *Config) AddFlags(flags *pflag.FlagSet) {
	flags.CountVarP(&c.Verbosity, "verbosity", "v", "log verbosity")
	flags.StringVarP(&c.File, "log-file", "", "", "log file")
	flags.BoolVarP(&c.Quiet, "quiet", "q", false, "quiet logs")
}

func (c *Config) PostLoad() error {
	// automatically configure the default logger after reading configuration
	return DefaultLogger(*c)
}

var _ config.PostLoad = (*Config)(nil)
