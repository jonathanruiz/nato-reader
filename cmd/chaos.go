/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chaosCmd represents the chaos command
var chaosCmd = &cobra.Command{
	Use:   "chaos",
	Short: "Print out a confusing phoenetic output of a word.",
	Long: `chaos will for print out a phonetic output of a word except it will be hard to understand. The ouput will be the most misleading and confusing words.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chaos called")
	},
}

func init() {
	rootCmd.AddCommand(chaosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chaosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chaosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
