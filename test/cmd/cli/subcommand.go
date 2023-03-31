package cli

import (
	"github.com/spf13/cobra"

	"github.com/anchore/go-cli-tools/config"
	"github.com/anchore/go-cli-tools/inject"
	"github.com/anchore/go-cli-tools/test/cmd/cli/options"
)

func Subcommand(c inject.Container, formatOpts *options.Format) *cobra.Command {
	cmd := &cobra.Command{
		Use: "subcommand",
		Args: func(cmd *cobra.Command, args []string) error {
			return config.Load(c, cmd, formatOpts)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	formatOpts.BindFlags(cmd.Flags())

	return cmd
}
