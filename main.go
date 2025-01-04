package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	var displayEquation bool
	nums := make([]int, 0, 2)

	var validator = func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("you must provide 2 arguments")
		}
		for _, arg := range args {
			num, err := strconv.Atoi(arg)
			if err != nil {
				return errors.New("only integer arguments are supported")
			}
			nums = append(nums, num)
		}
		if cmd.Use == "divide" {
			if nums[1] == 0 {
				return errors.New("divide by 0 error")
			}
		}
		return nil
	}

	// creating the different sub-commands
	var add = &cobra.Command{Use: "add", Args: validator}
	var multiply = &cobra.Command{Use: "multiply", Args: validator}
	var subtract = &cobra.Command{Use: "subtract", Args: validator}
	var divide = &cobra.Command{Use: "divide", Args: validator}

	// attaching a flag to all the sub-commands
	add.Flags().BoolVar(&displayEquation, "display-equation", false, "whether to display the complete equation")
	multiply.Flags().BoolVar(&displayEquation, "display-equation", false, "whether to display the complete equation")
	subtract.Flags().BoolVar(&displayEquation, "display-equation", false, "whether to display the complete equation")
	divide.Flags().BoolVar(&displayEquation, "display-equation", false, "whether to display the complete equation")

	// attaching callbacks to all the sub-commands
	add.Run = func(cmd *cobra.Command, args []string) {
		if displayEquation {
			fmt.Printf("%v + %v = %v\n", nums[0], nums[1], nums[0]+nums[1])
		} else {
			fmt.Println(nums[0] + nums[1])
		}
	}

	multiply.Run = func(cmd *cobra.Command, args []string) {
		if displayEquation {
			fmt.Printf("%v x %v = %v\n", nums[0], nums[1], nums[0]*nums[1])
		} else {
			fmt.Println(nums[0] * nums[1])
		}
	}

	subtract.Run = func(cmd *cobra.Command, args []string) {
		if displayEquation {
			fmt.Printf("%v - %v = %v\n", nums[0], nums[1], nums[0]-nums[1])
		} else {
			fmt.Println(nums[0] - nums[1])
		}
	}

	divide.Run = func(cmd *cobra.Command, args []string) {
		if displayEquation {
			fmt.Printf("%v / %v = %v\n", nums[0], nums[1], nums[0]/nums[1])
		} else {
			fmt.Println(nums[0] / nums[1])
		}
	}

	// initializing the root command
	var calculator = &cobra.Command{Use: "calculator"}

	// adding all the sub-commands to this root command
	calculator.AddCommand(add, multiply, subtract, divide)
	calculator.Execute()

}
