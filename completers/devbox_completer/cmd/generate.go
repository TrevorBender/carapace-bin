package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate supporting files for your project",
	Args:  cobra.MaximumNArgs(0),
}

func init() {
	carapace.Gen(genDocsCmd).Standalone()
	rootCmd.AddCommand(genDocsCmd)
	generateCmd.AddCommand(devcontainerCmd())
	generateCmd.AddCommand(dockerfileCmd())
	generateCmd.AddCommand(debugCmd())
	generateCmd.AddCommand(direnvCmd())
	generateCmd.AddCommand(sshConfigCmd())
}

func debugCmd() *cobra.Command {
	command := &cobra.Command{
		Use:    "debug",
		Hidden: true,
		Args:   cobra.MaximumNArgs(0),
		Run:    func(cmd *cobra.Command, args []string) {},
	}
	return command
}

func devcontainerCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "devcontainer",
		Short: "Generate Dockerfile and devcontainer.json files under .devcontainer/ directory",
		Long:  "Generate Dockerfile and devcontainer.json files necessary to run VSCode in remote container environments.",
		Args:  cobra.MaximumNArgs(0),
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	command.Flags().BoolP("force", "f", false, "force overwrite on existing files")
	return command
}

func dockerfileCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "dockerfile",
		Short: "Generate a Dockerfile that replicates devbox shell",
		Long: "Generate a Dockerfile that replicates devbox shell. " +
			"Can be used to run devbox shell environment in an OCI container.",
		Args: cobra.MaximumNArgs(0),
		Run:  func(cmd *cobra.Command, args []string) {},
	}
	command.Flags().BoolP("force", "f", false, "force overwrite existing files")
	return command
}

func direnvCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "direnv",
		Short: "Generate a .envrc file that integrates direnv with this devbox project",
		Long: "Generate a .envrc file that integrates direnv with this devbox project. " +
			"Requires direnv to be installed.",
		Args: cobra.MaximumNArgs(0),
		Run:  func(cmd *cobra.Command, args []string) {},
	}
	command.Flags().BoolP("force", "f", false, "force overwrite existing files")
	command.Flags().BoolP("print-envrc", "p", false, "output contents of devbox configuration to use in .envrc")
	// this command marks a flag as hidden. Error handling for it is not necessary.
	_ = command.Flags().MarkHidden("print-envrc")

	return command
}

func sshConfigCmd() *cobra.Command {
	command := &cobra.Command{
		Use:    "ssh-config",
		Hidden: true,
		Short:  "Generate ssh config to connect to devbox cloud",
		Long:   "Check ssh config and if they don't exist, it generates the configs necessary to connect to devbox cloud VMs.",
		Args:   cobra.MaximumNArgs(0),
		Run:    func(cmd *cobra.Command, args []string) {},
	}
	command.Flags().StringP("username", "u", "", "GitHub username to use for ssh")
	return command
}
