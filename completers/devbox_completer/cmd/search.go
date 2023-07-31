package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <pkg>",
	Short: "Search for nix packages",
	Args:  cobra.ExactArgs(1),
}

func init() {
	carapace.Gen(searchCmd).Standalone()
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().Bool("show-all", false, "show all available templates")
}
