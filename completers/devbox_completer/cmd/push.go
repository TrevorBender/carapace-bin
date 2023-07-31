package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use: "push <git-repo>",
	Short: "Push a [global] config. Leave empty to use jetpack cloud. Can " +
		"be a git repo for self storage.",
	Args: cobra.MaximumNArgs(1),
}

func init() {
	carapace.Gen(pushCmd).Standalone()
	rootCmd.AddCommand(pushCmd)
}
