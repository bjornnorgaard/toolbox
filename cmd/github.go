package cmd

import (
	"fmt"
	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:     "github",
	Aliases: []string{"gh"},
	Short:   "Uses GitHub CLI to interact with repos",
	Long:    `Use one of the nested commands to actually do something`,
	Run:     runGithub(),
}

func runGithub() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		currentRepository, err := gh.CurrentRepository()
		if err != nil {
			fmt.Printf("you are likely not in a git repo, anyway here is an error: %v", err)
			return
		}
		fmt.Printf("Signed into %s as %s", currentRepository.Host(), currentRepository.Owner())
	}
}

func init() {
	rootCmd.AddCommand(githubCmd)
}
