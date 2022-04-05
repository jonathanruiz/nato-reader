/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

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

		outputNatoWord(natoWord)
	},
}

// Struct for each Nato letter
type Letter struct {
	Letter string `json:"letter"`
	Word   string `json:"word"`
}

// Execute adds all child commands to the root command and sets flags appropriately.
func init() {
	rootCmd.AddCommand(wordCmd)

	// This flag is where the user adds the word they are wanting to get a phonetic output
	wordCmd.Flags().StringP("word", "w", "", "The word you are looking to get the phonetic output for")

}

// Converts the word to uppercase and creates an array of strings
func ConvertWord(word string) []string {

	// Create uppercase version of the word
	var uppercaseWord string = strings.ToUpper(word)

	// Create an array of strings
	return strings.Split(uppercaseWord, "")

}

// Outputs the word from the word flag
func outputNatoWord(word string) {
	// Create an array of empty Letter structs
	var letters []Letter

	// Open our jsonFile
	content, err := ioutil.ReadFile("./json/nato.json")

	// If we ioutil.ReadFile returns an error then handle it
	if err != nil {
		fmt.Println(err.Error())
	}

	// Unmarshal the jsonFile into the Letter struct array
	err2 := json.Unmarshal(content, &letters)

	// If we json.Unmarshal returns an error then handle it
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}

	// Convert word to uppercase and create an array of strings
	var wordArray = ConvertWord(word)

	// Loop through the wordArray
	for _, character := range wordArray {
		// Loop through the letters array
		for _, letter := range letters {
			// If the letter in the wordArray matches the letter in the letters array
			if letter.Letter == character {
				fmt.Println(letter.Letter + " as in " + letter.Word)
			}
		}
	}
}
