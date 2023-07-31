package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var integrateCmd = &cobra.Command{
	Use:    "integrate",
	Short:  "integrate with an IDE",
	Args:   cobra.MaximumNArgs(1),
	Hidden: true,
}

func init() {
	carapace.Gen(integrateCmd).Standalone()
	rootCmd.AddCommand(integrateCmd)
	integrateCmd.AddCommand(integrateVSCodeCmd())
}

func integrateVSCodeCmd() *cobra.Command {
	command := &cobra.Command{
		Use:    "vscode",
		Hidden: true,
		Short:  "Integrate devbox environment with VSCode.",
	}

	return command
}
