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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show the desired phonetic alphabet",
	Long: `This command will show the desired phonetic alphabet. Add the -a flag to select the specific alphabet.`,
	Run: func(cmd *cobra.Command, args []string) {
		alphabetFlag, _ := cmd.Flags().GetString("alpha")

		// If the alphabet flag is empty then output the entire alphabet. Otherwise, output the specific alphabet.
		if alphabetFlag == "chaos" {
			outputAlphabet("./json/chaos.json")
		} else {
			outputAlphabet("./json/nato.json")
		}

	},
}

func outputAlphabet(jsonFile string) {
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

	// Loop through the Letter struct array
	for _, letter := range letters {
		// Print the Letter struct values
		fmt.Println(letter.Letter + " as in " + letter.Word)
	}

}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("alpha", "a", "", "The specific alphabet you want to output")
}
