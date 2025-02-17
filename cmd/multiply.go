package cmd

import (
	"cli-calculator/utils"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var Multiply = &cobra.Command{Use: "multiply", Args: utils.ValidateOperands}

func init() {
	var displayEquation bool
	Multiply.Flags().BoolVarP(&displayEquation, "display-equation", "e", false, "set this flag if you want to see the entire equation")

	Multiply.Run = func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.Atoi(args[0])
		num2, _ := strconv.Atoi(args[1])

		if displayEquation {
			fmt.Printf("%v x %v = %v\n", num1, num2, num1*num2)
		} else {
			fmt.Println(num1 * num2)
		}
	}
}
