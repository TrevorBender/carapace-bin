package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [dir] --template <template>",
	Short: "Initialize a directory as a devbox project using a template",
	Args:  cobra.MaximumNArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()
	createCmd.Flags().StringP("template", "t", "", "template to initialize the project with")
	createCmd.Flags().Bool("show-all", false, "show all available templates")
}
