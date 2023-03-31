package convert

import (
	"github.com/anchore/go-cli-tools/test/cmd/app"
	"github.com/anchore/go-cli-tools/test/cmd/cli/options"
)

func PackageBasicConfig(i *options.InputSource) *app.PackageBasicConfig {
	return &app.PackageBasicConfig{
		InputSource: i.InputSource,
	}
}
