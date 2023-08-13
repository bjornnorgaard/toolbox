package cmd

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github"
	"github.com/spf13/cobra"
)

var ()

func init() {
	githubCmd.AddCommand(approveCmd)
	approveCmd.Flags().BoolVarP(&dryRun, "dry-run", "n", false, "only fetch and print intentions, no actions performed")
}

var (
	dryRun = false

	// approveCmd represents the runApprove command
	approveCmd = &cobra.Command{
		Use:     "approve-dependabot-pulls",
		Aliases: []string{"approve", "ad", "a"},
		Short:   "Approves pull-requests by dependabot",
		Long: "Approves pull-requests by dependabot.\n" +
			"Will only approve pull requests which are passing CI checks\n" +
			"and not already reviewed",
		Run: runApprove(),
	}
)

func runApprove() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := github.Approve(dryRun); err != nil {
			fmt.Printf("❌ Approve failed: %v", err)
		}
	}
}
