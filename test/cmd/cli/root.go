package cli

import (
	"github.com/spf13/cobra"

	"github.com/anchore/fangs"
	"github.com/anchore/go-cli-tools/log"
	"github.com/anchore/go-cli-tools/test/cmd/cli/options"
)

func Root(pkgs *cobra.Command) any {
	return func(cfg *fangs.Config, log *log.Config, commandConfig *options.CommandConfig) *cobra.Command {
		cmd := &cobra.Command{
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return fangs.LoadAt(*cfg, cmd, "log", log)
			},
			Args: pkgs.Args,
			RunE: pkgs.RunE,
		}

		flags := cmd.PersistentFlags()
		cfg.AddFlags(flags)
		log.AddFlags(flags)

		flags = cmd.Flags()
		commandConfig.AddFlags(flags)

		return cmd
	}
}
