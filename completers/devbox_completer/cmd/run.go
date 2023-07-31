package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [<script> | <cmd>]",
	Short: "Run a script or command in a shell with access to your packages",
	Long: "Start a new shell and runs your script or command in it, exiting when done.\n\n" +
		"The script must be defined in `devbox.json`, or else it will be interpreted as an " +
		"arbitrary command. You can pass arguments to your script or command. Everything " +
		"after `--` will be passed verbatim into your command (see examples).\n\n",
	Example: "\nRun a command directly:\n\n  devbox add cowsay\n  devbox run cowsay hello\n  " +
		"devbox run -- cowsay -d hello\n\nRun a script (defined as `\"moo\": \"cowsay moo\"`) " +
		"in your devbox.json:\n\n  devbox run moo",
}

func init() {
	carapace.Gen(runCmd).Standalone()
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().Bool("pure", false, "If this flag is specified, devbox runs the script in an isolated environment inheriting almost no variables from the current environment. A few variables, in particular HOME, USER and DISPLAY, are retained.")
	runCmd.Flags().BoolP("list", "l", false, "List all scripts defined in devbox.json")

	//runCmd.ValidArgs = listScripts(command)
}
