package cmd

import (
	"fmt"

	"github.com/bjornnorgaard/toolbox/tools/github"
	"github.com/spf13/cobra"
)

func init() {
	approveCmd.Flags().BoolP("force", "f", false, "Include already approved PRs for processing")
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
			"and not already reviewed.\n" +
			"Use --force/-f to include already approved PRs for processing.",
		Run: runApprove(),
	}
)

func runApprove() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		if err := github.Approve(force); err != nil {
			fmt.Println(err)
		}
	}
}
