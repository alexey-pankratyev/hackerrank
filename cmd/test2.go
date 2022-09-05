/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// test2Cmd represents the test2 command
var test2Cmd = &cobra.Command{
	Use:   "test2",
	Short: "test2",
	Long: `test2`,
	Run: func(cmd *cobra.Command, args []string) {
		results := runTest2()
		fmt.Printf("test2 results: %v \n", results)

	},
}

func runTest2() string {
	ch := rune(97)
	n := int('a')
	return fmt.Sprintf("\nchar: %c\ncode: %d\n", ch, n)
}

func init() {
	rootCmd.AddCommand(test2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
