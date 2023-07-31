package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var genDocsCmd = &cobra.Command{
	Use:   "gen-docs <path>",
	Short: "[Internal] Generate documentation for the CLI",
	Long: "[Internal] Generates the documentation for the CLI's Cobra commands. " +
		"Docs are placed in the directory specified by <path>.",
	Hidden: true,
	Args:   cobra.ExactArgs(1),
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genDocsCmd).Standalone()
	rootCmd.AddCommand(genDocsCmd)
}
