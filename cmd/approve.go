package cmd

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github"
	"github.com/spf13/cobra"
)

// approveCmd represents the approve command
var approveCmd = &cobra.Command{
	Use:     "approve",
	Aliases: []string{"a"},
	Short:   "Approves pull-requests by dependabot",
	Long:    `Will approved pull-requests made by dependabot which are passed CI and not previously reviewed`,
	Run:     runApprove(),
}

func runApprove() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := github.ApproveDependabotPullRequests()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func init() {
	githubCmd.AddCommand(approveCmd)
}
