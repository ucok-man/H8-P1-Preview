package types

import (
	"fmt"
)

type operation string

const (
	Add operation = "add"
	Sub operation = "sub"
	Mul operation = "mul"
	Div operation = "div"
)

type Calculator struct {
	Operation operation
	Num1      float64
	Num2      float64
	result    *float64
}

func (c *Calculator) Addition() {
	calculation := c.Num1 + c.Num2
	c.result = &calculation
}

func (c *Calculator) Subtraction() {
	calculation := c.Num1 - c.Num2
	c.result = &calculation
}

func (c *Calculator) Multiply() {
	calculation := c.Num1 * c.Num2
	c.result = &calculation
}

func (c *Calculator) Division() {
	calculation := c.Num1 / c.Num2
	c.result = &calculation
}

func (c *Calculator) GetResult() string {
	if nil == c.result {
		panic("Calling result without performing operation")
	}
	return fmt.Sprintf("Result\t: %.2f", *c.result)
}
