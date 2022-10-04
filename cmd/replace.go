package cmd

import (
	"fmt"

	"github.com/bjornnorgaard/toolbox/replace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(replaceCmd)
	replaceCmd.Flags().StringVarP(&char, "char", "c", "X", "the replacement char")
	replaceCmd.Flags().StringVarP(&method, "method", "m", "", "the method to print replaced")
}

var (
	method = ""
	char   = ""
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:     "replace <file name>",
	Aliases: []string{"r"},
	Args:    cobra.ExactArgs(1),
	Short:   "Can change all chars of either a method or file to a single char.",
	Run:     runReplace(),
	Long:    `Sometimes it's fun to find new ways to check the readability of files.`,
}

func runReplace() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		file := args[0]
		result, err := replace.ConvertToMonoChar(file, method, char)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}
