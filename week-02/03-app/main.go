package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ucok-man/PHASE_1/week_2/day1/03-app/types"
)

func main() {
	if err := run(os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(reader io.Reader) error {
	var app types.Application

	input, err := app.Pompt("Enter Number: ", reader)
	if err != nil {
		return err
	}

	app.Data, err = app.ValidateInput(input)
	if err != nil {
		return err
	}

	result := app.Process()
	fmt.Println(result)

	return nil
}
