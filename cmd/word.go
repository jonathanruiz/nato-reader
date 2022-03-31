/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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

func init() {
	rootCmd.AddCommand(wordCmd)

	// This flag is where the user adds the word they are wanting to get a phonetic output
	wordCmd.Flags().StringP("word", "w", "", "The word you are looking to get the phonetic output for")

}

// This outputs the word from the word flag
func outputNatoWord(word string) {
	// Object for each Nato letter
	type Letter struct {
		Letter string `json:"letter"`
		Word   string `json:"word"`
	}

	// Open our jsonFile
	content, err := ioutil.ReadFile("./json/nato.json")

	// If we ioutil.ReadFile returns an error then handle it
	if err != nil {
		fmt.Println(err.Error())
	}

	// Create an array of empty Letter structs
	var letters []Letter

	// Parse json into array of Letter structs
	err2 := json.Unmarshal(content, &letters)

	// Error handling again
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}

	for _, x := range letters {
		fmt.Println(x.Word)
	}

	for _, letter := range word {

		fmt.Println(string(letter))
	}
}
