package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// To be used for the 'add', 'subtract', 'multiply', 'divide' commands
func ValidateOperands(command *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("you must provide 2 arguments")
	}
	for _, arg := range args {
		_, err := strconv.Atoi(arg)
		if err != nil {
			return errors.New("only integer arguments are supported")
		}
	}
	if command.Use == "divide" {
		divisor, _ := strconv.Atoi(args[1])
		if divisor == 0 {
			return errors.New("divide by 0 error")
		}
	}
	return nil
}

// To be used for the 'evaluate' command
func ValidateExpression(command *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("you must provide 1 argument")
	}

	allowedChars := []rune{'(', ')', '+', '-', '*', '/', ' '}

	expression := args[0]
	for _, char := range expression {
		if contains(allowedChars, char) {
			continue
		}
		if _, err := strconv.Atoi(string(char)); err != nil {
			return fmt.Errorf("%c is an invalid character", char)
		}
	}
	return nil
}

func contains(charList []rune, char rune) bool {
	for _, val := range charList {
		if val == char {
			return true
		}
	}
	return false
}
