/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github"
	"github.com/spf13/cobra"
)

// enableAutoMergeCmd represents the enableAutoMerge command
var enableAutoMergeCmd = &cobra.Command{
	Use:     "enable-auto-merge",
	Aliases: []string{"eam"},
	Short:   "Enables auto-merge for all repos",
	// Long: ``,
	Run: runEnableAutoMerge(),
}

func runEnableAutoMerge() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := github.EnableAutoMerge()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func init() {
	githubCmd.AddCommand(enableAutoMergeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enableAutoMergeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enableAutoMergeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
