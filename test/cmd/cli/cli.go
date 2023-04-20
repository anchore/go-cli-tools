package cli

import (
	"github.com/spf13/cobra"

	"github.com/anchore/go-cli-tools/command"
	"github.com/anchore/go-cli-tools/config"
	"github.com/anchore/go-cli-tools/inject"
	"github.com/anchore/go-cli-tools/log"
	"github.com/anchore/go-cli-tools/test/cmd/cli/convert"
	"github.com/anchore/go-cli-tools/test/cmd/cli/options"
)

func New() (*cobra.Command, error) {
	c := inject.NewContainer()

	c.Bind(
		// global options Root adds to persistent flags
		config.NewConfig("app"),
		log.NewConfig(),

		// format options are shared among multiple commands
		options.NewFormat(),
	)

	c.Register(
		// command options are shared between Root and Command
		inject.Singleton(options.NewCommandConfig),
	)

	command.Converters(
		c,
		convert.CatalogerConfig,
		convert.GoCatalogerConfig,
	)

	pkgs := command.MakeOne(c, Command) // Packages is a snowflake we copy to the root command
	r := command.MakeOne(c, Root(pkgs)) // Root binds persistent root flags
	r.AddCommand(
		pkgs,
	)
	r.AddCommand(
		command.Make(c,
			Subcommand,
		)...,
	)

	return r, nil
}
