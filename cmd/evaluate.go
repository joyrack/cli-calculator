package cmd

import (
	"cli-calculator/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var Evaluate = &cobra.Command{Use: "evaluate", Args: utils.ValidateExpression}

func init() {
	Evaluate.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("Evaluate called")
	}
}
