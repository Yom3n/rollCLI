/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"strconv"
	"strings"

	"github.com/Yom3n/rollCLI/dice"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "roll",
	Short: "Roll the dice using xdy",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			splitInput := strings.Split(arg, "d")
			numDices, _ := strconv.Atoi(splitInput[0])
			sides, sidesErr := strconv.Atoi(splitInput[1])
			if numDices == 0 {
				// Handles input in shortened form, For exmaple "d6" instead of "1d6"
				numDices = 1
				if sidesErr != nil {
					fmt.Errorf("%v is invalid input. Input must be in format XdY. For example 1d6, or d6", arg)
				}
				var dice dice.Dice = dice.Dice()
				dice.SetSides(uint(sides))
				output := arg
				output += ": "
				for i := 0; i < numDices; i++ {
					roll := dice.Roll()
					output += roll
					if i != numDices-1 {
						output += ", "
					}
				}
				fmt.Println(output)
			}

		}
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rollCLI.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
