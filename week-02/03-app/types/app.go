package types

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"
)

type Application struct {
	Data int
}

func (app Application) Pompt(question string, reader io.Reader) (string, error) {
	fmt.Print(question)

	input, err := app.inputFromStdin(reader)
	if err != nil {
		return "", err
	}

	return input, nil

}

func (app Application) inputFromStdin(reader io.Reader) (string, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", nil
	}
	return scanner.Text(), nil
}

func (app Application) ValidateInput(input string) (int, error) {
	if input == "" {
		return -1, fmt.Errorf("invalid input: field is required")
	}

	inputNum, err := strconv.Atoi(input)
	if err != nil {
		return -1, fmt.Errorf("invalid input (%s): make sure entering valid number", input)
	}

	if inputNum < 2 {
		return -1, fmt.Errorf("invalid input (%d): must be positive >= 2", inputNum)
	}

	return inputNum, nil
}

func (app Application) Process() string {
	var buff bytes.Buffer

	buff.WriteString(fmt.Sprintf("\t\t\t\t%s\n", "~ RESULT ~"))
	buff.WriteString(app.evenOdd())
	buff.WriteString(app.primenumber())

	return buff.String()
}

func (app Application) evenOdd() string {
	var buff bytes.Buffer

	var iseven bool
	if app.Data%2 == 0 {
		buff.WriteString("your number is even.\n")
		iseven = true
	}
	if app.Data%2 != 0 {
		buff.WriteString("your number is odd.\n")
		iseven = false
	}

	var evenupto []int
	for i := 1; i <= app.Data; i++ {
		switch iseven {
		case true:
			if i%2 == 0 {
				evenupto = append(evenupto, i)
			}
			continue
		case false:
			if i%2 != 0 {
				evenupto = append(evenupto, i)
			}
		default:
			panic("unhandled posible case")
		}
	}
	buff.WriteString(fmt.Sprintf("even number upto your input: %v\n", evenupto))

	return buff.String()
}

func (app Application) primenumber() string {
	var buff bytes.Buffer

	sq_root := int(math.Sqrt(float64(app.Data)))
	for i := 2; i <= sq_root; i++ {
		if app.Data%i == 0 {
			buff.WriteString("your number is not a prime number\n")
			return buff.String()
		}
	}
	buff.WriteString("your number is a prime number\n")
	return buff.String()
}
