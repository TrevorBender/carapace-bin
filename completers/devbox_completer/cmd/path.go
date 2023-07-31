package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show path to [global] devbox config",
	Args:  cobra.ExactArgs(0),
}

func init() {
	carapace.Gen(pathCmd).Standalone()
	rootCmd.AddCommand(pathCmd)
}
