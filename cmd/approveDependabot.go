package cmd

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github"
	"github.com/spf13/cobra"
)

// approveCmd represents the approve command
var approveCmd = &cobra.Command{
	Use:     "approve-dependabot",
	Aliases: []string{"ad"},
	Short:   "Approves pull-requests by dependabot",
	Long: `Approves pull-requests by dependabot
	Only pull-requests which are passing CI 
	and not already approved in review`,
	Run: runApprove(),
}

func runApprove() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := github.ApproveDependabotPullRequests(dryRun)
		if err != nil {
			fmt.Println(err)
		}
	}
}

var (
	dryRun = false
)

func init() {
	githubCmd.AddCommand(approveCmd)
	approveCmd.Flags().BoolVarP(&dryRun, "dry-run", "n", false, "only fetch and print intentions, no actions performed")
}
