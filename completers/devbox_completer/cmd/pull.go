package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull <file> | <url>",
	Short: "Pull a config from a file or URL",
	Long:  "Pull a config from a file or URL. URLs must be prefixed with 'http://' or 'https://'.",
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	carapace.Gen(pullCmd).Standalone()
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().BoolP("force", "f", false, "Force overwrite of existing [global] config files")
}
