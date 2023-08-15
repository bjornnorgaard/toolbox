package cmd

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(githubCmd)
}

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:     "github",
	Aliases: []string{"gh", "g"},
	Short:   "Uses GitHub CLI to interact with repos",
	Long:    `Use one of the nested commands to actually do something`,
	Run:     runGithub(),
}

func runGithub() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := github.Github()
		if err != nil {
			fmt.Println(err)
		}
	}
}
