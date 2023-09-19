package cmd

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github"
	"github.com/spf13/cobra"
)

func init() {
	githubCmd.AddCommand(automergeCmd)
}

// automergeCmd represents the automerge command
var automergeCmd = &cobra.Command{
	Use:     "automerge",
	Aliases: []string{"automerge", "am"},
	Short:   "Sets specific pull requests to auto merge",
	Long: `Sets all pull requests that fit criteria to auto merge. So far these criteria are:
	- Pull request is from dependabot
	- Pull request is passing CI checks
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := github.SetAutoMerge()
		if err != nil {
			fmt.Println(err)
		}
	},
}
