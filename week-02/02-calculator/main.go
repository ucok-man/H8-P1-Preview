package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ucok-man/PHASE_1/week_2/day1/02-calculator/types"
)

func main() {
	if len(os.Args[1:]) != 3 {
		showUsage(os.Stderr)
		os.Exit(1)
	}

	if os.Args[1] == "help" {
		showUsage(os.Stderr)
		os.Exit(1)
	}

	param := types.Input{
		Operation: strings.ToLower(os.Args[1]),
		Operand_1: strings.ToLower(os.Args[2]),
		Operand_2: strings.ToLower(os.Args[3]),
	}
	if err := run(param); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func run(input types.Input) error {
	if err := input.Validate(); err != nil {
		return err
	}

	app := input.ToCalculator()

	switch app.Operation {
	case types.Add:
		app.Addition()
	case types.Sub:
		app.Subtraction()
	case types.Mul:
		app.Multiply()
	case types.Div:
		if app.Num2 == 0 {
			return fmt.Errorf("invalid input: zero division")
		}
		app.Division()
	default:
		panic(fmt.Sprintf("undhadled posible case: %v", os.Args[1]))
	}

	fmt.Println(app.GetResult())
	return nil
}

func showUsage(writer io.Writer) {
	fmt.Fprintln(writer, "~calculator~")
	fmt.Fprintln(writer, "Usage: calculator [help] | [<operation> <operand> <operand>]")
}
