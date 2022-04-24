/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Struct for each Nato letter
type Letter struct {
	Letter string `json:"letter"`
	Word   string `json:"word"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nato-reader",
	Short: "Print out a phoenetic output of a word or sentence.",
	Long: `Nato-Reader is a CLI tool for printing out a phonetic output of a word or sentence.
	With this application, you can output words that would best fit your country.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Converts the word to uppercase and creates an array of strings
func convertWord(word string) []string {

	// Create uppercase version of the word
	var uppercaseWord string = strings.ToUpper(word)

	// Create an array of strings
	return strings.Split(uppercaseWord, "")

}

// Outputs the word from the word flag
func outputNatoWord(word string, jsonFile string) {
	// Create an array of empty Letter structs
	var letters []Letter

	// Open our jsonFile
	content, err := ioutil.ReadFile(jsonFile)

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
	var wordArray = convertWord(word)

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

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nato-reader.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}




