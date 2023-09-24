package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

// version, commit, and date are populated during build
var (
	version = "latest"
	commit  = "n/a"
	date    = "n/a"
)

func VersionCommand(errorCh *chan error) *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Conveyance version: %s, commit: %s, built at: %s", version, commit, date)
			*errorCh <- nil
		},
	}
}
