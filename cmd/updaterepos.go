package cmd

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github"
	"github.com/spf13/cobra"
)

func init() {
	githubCmd.AddCommand(updateReposCmd)
}

var (
	// updateReposCmd represents the enableAutoMerge command
	updateReposCmd = &cobra.Command{
		Use:     "update-repos",
		Aliases: []string{"ur"},
		Short:   "Enables auto-merge and some other stuff for all repos owner by current user",
		Run:     runUpdateRepos(),
	}
)

func runUpdateRepos() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := github.UpdateRepos()
		if err != nil {
			fmt.Printf("❌ Failed to update repos: %v", err)
		}
	}
}
