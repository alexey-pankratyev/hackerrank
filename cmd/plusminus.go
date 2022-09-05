/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

// plusminusCmd represents the plusminus command
var plusminusCmd = &cobra.Command{
	Use:   "plusminus",
	Short: "plusminus",
	Long: `plusminus`,
	Run: func(cmd *cobra.Command, args []string) {
		arrTemp := []int32{-4,3,-9,0,4,1}
		results := runPlusMinus(arrTemp)
		fmt.Printf("plusminus results:\n%v\n", results)
	},
}

func runPlusMinus(arr []int32) string{
    // Write your code here
	p := 0
	n := 0
	z := 0
	for i := range arr {
		switch{
			case math.Signbit(float64(arr[i])):
				n++
			case !math.Signbit(float64(arr[i])) && arr[i] != 0:
				p++
			case arr[i] == 0:
				z++  
		}				
	}
	pResult := fmt.Sprintf("%.6f\n",float64(p)/float64(len(arr)))
	nResult := fmt.Sprintf("%.6f\n",float64(n)/float64(len(arr)))
	zResult := fmt.Sprintf("%.6f\n",float64(z)/float64(len(arr)))
	return fmt.Sprintf("Positive: %vNegative: %vZero: %v",pResult,nResult,zResult)
}

func init() {
	rootCmd.AddCommand(plusminusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// plusminusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// plusminusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
