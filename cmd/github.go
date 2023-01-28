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
		repository, err := gh.CurrentRepository()
		if err != nil {
			fmt.Printf("unable to resolve current repo, %v", err)
			return
		}
		fmt.Println(repository)
	}
}

func init() {
	rootCmd.AddCommand(githubCmd)
}
