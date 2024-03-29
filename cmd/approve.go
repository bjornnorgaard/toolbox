package cmd

import (
	"fmt"

	"github.com/bjornnorgaard/toolbox/tools/github"
	"github.com/spf13/cobra"
)

func init() {
	githubCmd.AddCommand(approveCmd)
}

var (
	// approveCmd represents the runApprove command
	approveCmd = &cobra.Command{
		Use:     "approve-dependabot",
		Aliases: []string{"approve", "ad", "a"},
		Short:   "Approves dependabot pull requests",
		Long: "Approves pull-requests by dependabot.\n" +
			"Will only approve pull requests which are passing CI checks\n" +
			"and not already reviewed",
		Run: runApprove(),
	}
)

func runApprove() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := github.Approve(); err != nil {
			fmt.Println(err)
		}
	}
}
