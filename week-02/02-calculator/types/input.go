package types

import (
	"fmt"
	"strconv"
)

type Input struct {
	Operation string
	Operand_1 string
	Operand_2 string
}

func (i Input) Validate() error {
	safeoperation := []string{"add", "sub", "mul", "div"}

	var validoperation bool
	for _, sop := range safeoperation {
		if sop == i.Operation {
			validoperation = true
		}
	}
	if !validoperation {
		return fmt.Errorf("invalid operation (%s): permitted operation %v", i.Operation, safeoperation)
	}

	_, err := strconv.ParseFloat(i.Operand_1, 64)
	if err != nil {
		return fmt.Errorf("invalid operand (%s): make sure input correct number", i.Operand_1)
	}

	_, err = strconv.ParseFloat(i.Operand_2, 64)
	if err != nil {
		return fmt.Errorf("invalid operand (%s): make sure input correct number", i.Operand_2)
	}

	return nil
}

func (i Input) ToCalculator() *Calculator {
	var calc Calculator
	calc.Operation = operation(i.Operation)

	num1, _ := strconv.ParseFloat(i.Operand_1, 64)
	calc.Num1 = num1

	num2, _ := strconv.ParseFloat(i.Operand_2, 64)
	calc.Num2 = num2

	return &calc
}
