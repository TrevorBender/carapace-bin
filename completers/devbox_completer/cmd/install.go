package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install all packages mentioned in devbox.json",
	Args:  cobra.MaximumNArgs(0),
}

func init() {
	carapace.Gen(installCmd).Standalone()
	rootCmd.AddCommand(installCmd)
}
