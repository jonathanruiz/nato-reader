/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> b69aa4b4b9447304aa3e8d0c38e110af567397ab
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nato-reader",
	Short: "Print out a phoenetic output of a word.",
	Long: `Nato-Reader is a CLI tool for printing out a phonetic output of a word.
	With this application, you can output words that would best fit your country.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
<<<<<<< HEAD
	Run: func(cmd *cobra.Command, args []string) { fmt.Println("Hello Nato-Reader") },
=======
	// Run: func(cmd *cobra.Command, args []string) { },
>>>>>>> b69aa4b4b9447304aa3e8d0c38e110af567397ab
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
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
