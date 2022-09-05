/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long: `test`,
	Run: func(cmd *cobra.Command, args []string) {
		results := runTest()
		fmt.Printf("test results: %v \n", results)

	},
}

func runTest() string {
	arr0 := []string{"men","woman", "child", "woman"}
	map0 := map[int]string{}
	mapRes := map[string]int{}
	arrResult := []string{}
	for i:=0; i<len(arr0);i++ {
		map0[i] = arr0[i]
	} 	
	for k, v := range map0 {
		if _, ok := mapRes[v]; ok {
			delete(map0,k)
		}else{
			mapRes[v] = k
			arrResult = append(arrResult,v) 	
		}
	}
	return fmt.Sprintf("\n%v\n",arrResult)
	
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
}
