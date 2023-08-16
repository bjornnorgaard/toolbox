package cmd

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github"
	"github.com/spf13/cobra"
)

func init() {
	githubCmd.AddCommand(updateReposCmd)
	updateReposCmd.Flags().BoolVarP(&dryRun, "dry-run", "n", false, "only fetch and print intentions, no write actions performed")

}

var (
	// Flags
	dryRun bool

	// updateReposCmd represents the enableAutoMerge command
	updateReposCmd = &cobra.Command{
		Use:     "update-repos",
		Aliases: []string{"ur"},
		Short:   "Enables repo settings I like",
		Long:    "Enables auto-merge, squash merge, branch update on PR and automatic branch deletion after merge",
		Run:     runUpdateRepos(),
	}
)

func runUpdateRepos() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := github.UpdateRepos(dryRun)
		if err != nil {
			fmt.Println(err)
		}
	}
}
