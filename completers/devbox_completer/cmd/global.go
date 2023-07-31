package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var globalCmd = &cobra.Command{
	Use:   "global",
	Short: "Manage global devbox packages",
}

func init() {
	carapace.Gen(globalCmd).Standalone()
	rootCmd.AddCommand(globalCmd)

	//addCommandAndHideConfigFlag(globalCmd, addCmd())
	//addCommandAndHideConfigFlag(globalCmd, installCmd())
	//addCommandAndHideConfigFlag(globalCmd, pathCmd())
	//addCommandAndHideConfigFlag(globalCmd, pullCmd())
	//addCommandAndHideConfigFlag(globalCmd, pushCmd())
	//addCommandAndHideConfigFlag(globalCmd, removeCmd())
	//addCommandAndHideConfigFlag(globalCmd, runCmd())
	//addCommandAndHideConfigFlag(globalCmd, servicesCmd())
	//addCommandAndHideConfigFlag(globalCmd, shellEnvCmd())
	//addCommandAndHideConfigFlag(globalCmd, updateCmd())

	// Create list for non-global? Mike: I want it :)
	globalCmd.AddCommand(globalListCmd())
}

func globalListCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List global packages",
	}
}
