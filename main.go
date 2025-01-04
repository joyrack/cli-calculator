package main

import (
	"cli-calculator/cmd"

	"github.com/spf13/cobra"
)

func main() {
	// initializing the root command
	var calculator = &cobra.Command{Use: "calculator"}

	// adding all the sub-commands to this root command
	calculator.AddCommand(cmd.Add, cmd.Multiply, cmd.Subtract, cmd.Divide, cmd.Evaluate)
	calculator.Execute()
}
