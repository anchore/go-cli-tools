package config

import (
	"os"
	"path"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/anchore/go-cli-tools/inject"
	"github.com/anchore/go-cli-tools/log"
	"github.com/anchore/go-logger/adapter/discard"
)

type sub struct {
	Sv      string `json:"sv" yaml:"sv" mapstructure:"sv"`
	Unbound string `json:"unbound" yaml:"unbound" mapstructure:"unbound"`
}

type root struct {
	Log *log.Config `json:"log" yaml:"log" mapstructure:"log"`
	V   string      `json:"v" yaml:"v" mapstructure:"v"`
	Sub *sub        `json:"sub" yaml:"sub" mapstructure:"sub"`
}

func Test_LoadDefaults(t *testing.T) {
	c, cmd, _, r, s, lc := setup(t)

	err := Load(c, cmd, lc, r)
	require.NoError(t, err)

	require.Equal(t, "default-sv", s.Sv)
	require.Equal(t, "default-v", r.V)
}

func Test_LoadFromConfigFile(t *testing.T) {
	c, cmd, cfg, r, s, _ := setup(t)

	wd, err := os.Getwd()
	require.NoError(t, err)
	cfg.ConfigFile = path.Join(wd, "test-fixtures", "config.yaml")

	err = Load(c, cmd, r)
	require.NoError(t, err)

	require.Equal(t, "config-sub-v", s.Sv)
	require.Equal(t, "config-v", r.V)
}

func Test_LoadFromEnv(t *testing.T) {
	t.Setenv("MY_APP_V", "env-var-v")
	t.Setenv("MY_APP_SUB_SV", "env-var-sv")

	c, cmd, _, r, s, _ := setup(t)

	err := Load(c, cmd, r)
	require.NoError(t, err)

	require.Equal(t, "env-var-sv", s.Sv)
	require.Equal(t, "env-var-v", r.V)
}

func Test_LoadFromEnvOverridingConfigFile(t *testing.T) {
	t.Setenv("MY_APP_V", "env-var-v")
	t.Setenv("MY_APP_SUB_SV", "env-var-sv")

	c, cmd, cfg, r, s, _ := setup(t)

	wd, err := os.Getwd()
	require.NoError(t, err)
	cfg.ConfigFile = path.Join(wd, "test-fixtures", "config.yaml")

	err = Load(c, cmd, r)
	require.NoError(t, err)

	require.Equal(t, "env-var-sv", s.Sv)
	require.Equal(t, "env-var-v", r.V)
}

func Test_LoadSubStruct(t *testing.T) {
	t.Setenv("MY_APP_SUB_SV", "env-var-sv")

	c, cmd, cfg, _, s, _ := setup(t)

	wd, err := os.Getwd()
	require.NoError(t, err)
	cfg.ConfigFile = path.Join(wd, "test-fixtures", "config.yaml")

	err = LoadAt(c, cmd, "sub", s)
	require.NoError(t, err)

	require.Equal(t, "env-var-sv", s.Sv)
}

func Test_LoadSubStructEnv(t *testing.T) {
	c, cmd, cfg, _, s, _ := setup(t)

	wd, err := os.Getwd()
	require.NoError(t, err)
	cfg.ConfigFile = path.Join(wd, "test-fixtures", "config.yaml")

	err = LoadAt(c, cmd, "sub", s)
	require.NoError(t, err)

	require.Equal(t, "config-sub-v", s.Sv)
}

func Test_LoadFromFlags(t *testing.T) {
	c, cmd, _, r, s, _ := setup(t)

	err := cmd.PersistentFlags().Set("v", "flag-value-v")
	require.NoError(t, err)

	err = cmd.Flags().Set("sv", "flag-value-sv")
	require.NoError(t, err)

	err = Load(c, cmd, r)
	require.NoError(t, err)

	require.Equal(t, "flag-value-sv", s.Sv)
	require.Equal(t, "flag-value-v", r.V)
}

func Test_LoadFromFlagsOverridingAll(t *testing.T) {
	t.Setenv("MY_APP_V", "env-var-v")
	t.Setenv("MY_APP_SUB_SV", "env-var-sv")

	c, cmd, cfg, r, s, _ := setup(t)

	wd, err := os.Getwd()
	require.NoError(t, err)
	cfg.ConfigFile = path.Join(wd, "test-fixtures", "config.yaml")

	err = cmd.PersistentFlags().Set("v", "flag-value-v")
	require.NoError(t, err)

	err = cmd.Flags().Set("sv", "flag-value-sv")
	require.NoError(t, err)

	err = Load(c, cmd, r)
	require.NoError(t, err)

	require.Equal(t, "flag-value-sv", s.Sv)
	require.Equal(t, "flag-value-v", r.V)
}

func setup(t *testing.T) (inject.Container, *cobra.Command, *Config, *root, *sub, *log.Config) {
	c := inject.NewContainer()

	cfg := NewConfig("my-app")
	c.Bind(cfg)

	s := &sub{
		Sv:      "default-sv",
		Unbound: "default-unbound",
	}

	logCfg := log.NewConfig()
	logCfg.Level = log.TraceLevel
	logCfg.Verbosity = 6

	t.Cleanup(func() {
		log.SetLogger(discard.New())
	})

	r := &root{
		Log: logCfg,
		V:   "default-v",
		Sub: s,
	}

	cmd := &cobra.Command{}

	flags := cmd.PersistentFlags()
	flags.StringVarP(&r.V, "v", "", r.V, "v usage")

	flags = cmd.Flags()
	flags.StringVarP(&s.Sv, "sv", "", s.Sv, "sv usage")

	return c, cmd, cfg, r, s, logCfg
}
