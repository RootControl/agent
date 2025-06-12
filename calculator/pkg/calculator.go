package pkg

import (
	"fmt"
	"strings"
)

type Calculator struct {
	Operators  map[string]func(float64, float64) float64
	Precedence map[string]int
}

func NewCalculator() *Calculator {
	return &Calculator{
		Operators: map[string]func(float64, float64) float64{
			"+": func(a, b float64) float64 { return a + b },
			"-": func(a, b float64) float64 { return a - b },
			"*": func(a, b float64) float64 { return a * b },
			"/": func(a, b float64) float64 { return a / b },
		},
		Precedence: map[string]int{
			"+": 1,
			"-": 1,
			"*": 2,
			"/": 2,
		},
	}
}

func (c *Calculator) Evaluate(expression string) (float64, error) {
	if expression == "" || expression == " " {
		return 0, fmt.Errorf("invalid expression")
	}

	tokens := strings.Fields(expression)
	return c.evaluateInfix(tokens)
}

func (c *Calculator) evaluateInfix(tokens []string) (float64, error) {
	var values []*float64
	var operators []string

	for _, token := range tokens {

	}
}

func (c *Calculator) applyOperator(operators *[]string, values *[]float64) error {
	if len(*operators) == 0 {
		return fmt.Errorf("no operators found")
	}

	operator := (*operators)[len(*operators)-1]
	*operators = (*operators)[:len(*operators)-1]

	if len(*values) < 2 {
		return fmt.Errorf("not enough operands for operator %s", operator)
	}

	b := (*values)[len(*values)-1]
	a := (*values)[len(*values)-2]
	*values = (*values)[:len(*values)-2]

	result := c.Operators[operator](a, b)
	*values = append(*values, result)
	return nil
}
