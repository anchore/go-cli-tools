package cli

import (
	"github.com/spf13/cobra"

	"github.com/anchore/fangs/config"
	"github.com/anchore/go-cli-tools/test/cmd/cli/options"
)

func Subcommand(cfg config.Config, formatOpts *options.Format) *cobra.Command {
	cmd := &cobra.Command{
		Use: "subcommand",
		Args: func(cmd *cobra.Command, args []string) error {
			return config.Load(cfg, cmd, formatOpts)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	formatOpts.AddFlags(cmd.Flags())

	return cmd
}
