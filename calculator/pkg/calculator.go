package pkg

import (
	"errors"
	"fmt"
	"strconv"
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
	var values []float64
	var operators []string

	for _, token := range tokens {
		if _, isOperator := c.Operators[token]; isOperator {
			for len(operators) > 0 &&
				c.Operators[operators[len(operators)-1]] != nil &&
				c.Precedence[operators[len(operators)-1]] >= c.Precedence[token] {
				if err := c.applyOperator(&operators, &values); err != nil {
					return 0, err
				}
			}

			operators = append(operators, token)
		} else {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid token: %s", token)
			}

			values = append(values, num)
		}
	}

	for len(operators) > 0 {
		if err := c.applyOperator(&operators, &values); err != nil {
			return 0, err
		}
	}

	if len(values) != 1 {
		return 0, errors.New("invalid expression")
	}

	return values[0], nil
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
