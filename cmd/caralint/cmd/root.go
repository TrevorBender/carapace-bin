package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "caralint",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exitCode := 0
		for _, arg := range args {
			if err := lint(arg); err != nil {
				println(err.Error())
				exitCode = 1
			}
		}
		os.Exit(exitCode)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".go"),
	)
}

func lint(path string) error {
	if !strings.HasSuffix(path, ".go") {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	r := regexp.MustCompile(`\.Flags\(\)\.(Bool|String)[^(]*\("(?P<name>[^"]+)"`)

	previous := ""
	line := 0
	for scanner.Scan() {
		if !r.MatchString(scanner.Text()) {
			previous = ""
		} else {
			matches := r.FindStringSubmatch(scanner.Text())
			current := matches[2]
			if previous != "" && previous > current {
				return fmt.Errorf("%v [%v]: flag '%v' should be before '%v'", path, line, current, previous)
			}
			previous = current
		}
		line += 1
	}
	return nil
}