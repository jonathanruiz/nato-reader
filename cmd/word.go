/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// wordCmd represents the word command
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "Print out a phoenetic output of a word",
	Long: `word will for print out a phonetic output of a word.
	With this application, you can output words that would best fit your country.`,
	Run: func(cmd *cobra.Command, args []string) {
		natoWord, _ := cmd.Flags().GetString("word")

		var natoJson string = normalAlphabet

		outputNatoWord(natoWord, natoJson)
	},
}


// Execute adds all child commands to the root command and sets flags appropriately.
func init() {
	rootCmd.AddCommand(wordCmd)

	// This flag is where the user adds the word they are wanting to get a phonetic output
	wordCmd.Flags().StringP("word", "w", "", "The word you are looking to get the phonetic output for")

}

