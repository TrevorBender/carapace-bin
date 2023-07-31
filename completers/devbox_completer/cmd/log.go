package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:    "log <event-name> [<event-specific-args>]",
	Hidden: true,
}

func init() {
	carapace.Gen(logCmd).Standalone()
	rootCmd.AddCommand(logCmd)
}
