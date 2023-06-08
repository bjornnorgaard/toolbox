package cmd

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/zeropad"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(zeropadCmd)
}

// zeropadCmd represents the zeropad command
var zeropadCmd = &cobra.Command{
	Use:     "zeropad",
	Aliases: []string{"zp"},
	Args:    cobra.ExactArgs(1),
	Run:     runZeroPad(),
	Short:   "Adds zero padding to image files in any nested folder",
	Long:    ``,
}

func runZeroPad() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := zeropad.CurrentDir(args[0])
		if err != nil {
			fmt.Println(err)
		}
	}
}
