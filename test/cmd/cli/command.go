package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/anchore/fangs/config"
	"github.com/anchore/go-cli-tools/inject"
	"github.com/anchore/go-cli-tools/test/cmd/app"
	"github.com/anchore/go-cli-tools/test/cmd/cli/options"
)

func Command(c inject.Container, cfg config.Config, formatOpts *options.Format) *cobra.Command {
	cmd := &cobra.Command{
		Use: "command",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("require an arg")
			}
			return config.Load(cfg, cmd, formatOpts)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := c.Invoke(app.Packages, args[0])
			return err
		},
	}

	formatOpts.AddFlags(cmd.Flags())

	return cmd
}
