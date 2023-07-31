package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:    "setup",
	Short:  "Setup devbox dependencies",
	Hidden: true,
}

func init() {
	carapace.Gen(setupCmd).Standalone()
	rootCmd.AddCommand(setupCmd)

	installNixCommand := &cobra.Command{
		Use:   "nix",
		Short: "Install Nix",
	}

	installNixCommand.Flags().Bool("daemon", false, "Install Nix in multi-user mode.")
	setupCmd.AddCommand(installNixCommand)
}
