package convert

import (
	"github.com/anchore/go-cli-tools/test/cmd/app"
	"github.com/anchore/go-cli-tools/test/cmd/cli/options"
)

func CatalogerConfig(c *options.CommandConfig) app.CatalogerConfig {
	cfg := app.DefaultCatalogerConfig()
	return cfg
}
