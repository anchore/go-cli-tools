package cli

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/anchore/go-cli-tools/config"
	"github.com/anchore/go-cli-tools/inject"
	"github.com/anchore/go-cli-tools/test/cmd/app"
	"github.com/anchore/go-cli-tools/test/cmd/cli/options"
)

func Command(c inject.Container, formatOpts *options.Format) *cobra.Command {
	cmd := &cobra.Command{
		Use: "command",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("require an arg")
			}
			c.Bind(options.InputSource{
				InputSource: args[0],
			})
			return config.Load(c, cmd, formatOpts)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Invoke(app.Packages)
		},
	}

	formatOpts.BindFlags(cmd.Flags())

	return cmd
}
