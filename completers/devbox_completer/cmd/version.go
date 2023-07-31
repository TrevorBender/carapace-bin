package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Args:  cobra.NoArgs,
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolP("verbose", "v", false, "displays additional version information")

	// Make this flag hidden because:
	// This functionality doesn't strictly belong in this versionCmd, but we add it here
	// since `devbox version update` calls `devbox version -v` to trigger an update.
	versionCmd.Flags().BoolP("update-devbox-symlink", "u", false, // value
		"update the devbox symlink to point to the current binary",
	)
	_ = versionCmd.Flags().MarkHidden("update-devbox-symlink")

	versionCmd.AddCommand(selfUpdateCmd())
}

func selfUpdateCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "update",
		Short: "Update devbox launcher and binary",
		Args:  cobra.ExactArgs(0),
	}

	return command
}
