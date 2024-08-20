/*
Copyright © 2024 Yomen <KamilBudnik94@gmail.com>
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

var helpStr = `	Roll is a CLI tool that let you quickly generate any dice rolls.
	User "roll XdY" to roll x y-sided dices.
	For exampel "roll 2d6" rolls 2 6-sided dices
	Roll accepts multiple arguments for example you can use "roll d20 2d6" to roll one d20 and two d6`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "roll",
	Short: "Roll the dice using xdy, where x is number of dices, and y is number of sides",
	Long:  helpStr,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(helpStr)
			return
		}

		for _, arg := range args {
			splitInput := strings.Split(arg, "d")
			numDices, numDicesErr := strconv.Atoi(splitInput[0])
			isNumDicesEmpty := splitInput[0] == ""
			sides, sidesErr := strconv.Atoi(splitInput[1])
			if numDices == 0 && isNumDicesEmpty {
				// Handles input in shortened form, For exmaple "d6" instead of "1d6"
				numDices = 1
			}
			if sidesErr != nil || (numDicesErr != nil && !isNumDicesEmpty) {
				fmt.Printf("\"%v\" is invalid input. Input must be in format XdY. For example 1d6, or d6", arg)
				return
			}
			var dice dice.Dice = dice.Dice{}
			dice.SetSides(uint(sides))
			output := generateFormattedOutput(dice, numDices, arg)
			fmt.Println(output)
		}
	},
}

// / Generates roll output in format:
// / [Input]: [rollResult1], [rollResult2] ... [rollResultN]
// / example: 2d6: 2, 6
func generateFormattedOutput(dice dice.Dice, numDices int, input string) string {
	output := input
	output += ": "
	for i := 0; i < numDices; i++ {
		roll := dice.Roll()
		output += strconv.Itoa(int(roll))
		if i != numDices-1 {
			output += ", "
		}
	}
	return output
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
