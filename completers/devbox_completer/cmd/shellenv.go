package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var shellEnvCmd = &cobra.Command{
	Use:   "shellenv",
	Short: "Print shell commands that add Devbox packages to your PATH",
	Args:  cobra.ExactArgs(0),
}

func init() {
	carapace.Gen(shellEnvCmd).Standalone()
	rootCmd.AddCommand(shellEnvCmd)
	shellEnvCmd.Flags().Bool("init-hook", false, "runs init hook after exporting shell environment")
	shellEnvCmd.Flags().Bool("install", false, "install packages before exporting shell environment")

	shellEnvCmd.Flags().Bool("pure", false, "If this flag is specified, devbox creates an isolated environment inheriting almost no variables from the current environment. A few variables, in particular HOME, USER and DISPLAY, are retained.")

	shellEnvCmd.AddCommand(shellEnvOnlyPathWithoutWrappersCmd())
}

func shellEnvOnlyPathWithoutWrappersCmd() *cobra.Command {
	command := &cobra.Command{
		Use:    "only-path-without-wrappers",
		Hidden: true,
		Short:  "[internal] Print shell command that exports the system $PATH without the bin-wrappers paths.",
		Args:   cobra.ExactArgs(0),
	}
	return command
}
