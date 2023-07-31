package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start a new shell with access to your packages",
	Long: "Start a new shell with access to your packages.\n\n" +
		"If the --config flag is set, the shell will be started using the devbox.json found in the --config flag directory. " +
		"If --config isn't set, then devbox recursively searches the current directory and its parents.",
	Args: cobra.NoArgs,
}

func init() {
	carapace.Gen(shellCmd).Standalone()
	shellCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")

	shellCmd.Flags().Bool("print-env", false, "print script to setup shell environment")
	shellCmd.Flags().Bool("pure", false, "If this flag is specified, devbox creates an isolated shell inheriting almost no variables from the current environment. A few variables, in particular HOME, USER and DISPLAY, are retained.")
	rootCmd.AddCommand(shellCmd)
	carapace.Gen(shellCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})
	carapace.Gen(shellCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
