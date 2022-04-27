/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// chaosCmd represents the chaos command
var chaosCmd = &cobra.Command{
	Use:   "chaos",
	Short: "Print out a confusing phoenetic output of a word.",
	Long: `chaos will for print out a phonetic output of a word except it will be hard to understand. The ouput will be the most misleading and confusing words.`,
	Run: func(cmd *cobra.Command, args []string) {
		natoWord, _ := cmd.Flags().GetString("word")

		var natoJson string = chaosAlphabet

		outputNatoWord(natoWord, natoJson)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func init() {
	rootCmd.AddCommand(chaosCmd)

	// This flag is where the user adds the word they are wanting to get a phonetic output
	chaosCmd.Flags().StringP("word", "w", "", "The word you are looking to get the phonetic output for")
}
