package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <pkg>",
	Short: "Display package info",
	Args:  cobra.ExactArgs(1),
}

func init() {
	carapace.Gen(infoCmd).Standalone()
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().Bool("markdown", false, "output in markdown format")
}
