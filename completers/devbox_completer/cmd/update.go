package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [pkg]...",
	Short: "Update packages in your devbox",
	Long: "Update one, many, or all packages in your devbox. " +
		"If no packages are specified, all packages will be updated. " +
		"Legacy non-versioned packages will be converted to @latest versioned " +
		"packages resolved to their current version.",
}

func init() {
	carapace.Gen(updateCmd).Standalone()
	rootCmd.AddCommand(updateCmd)
}
