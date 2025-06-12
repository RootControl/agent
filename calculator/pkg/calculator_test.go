package pkg

import "testing"

func TestEvaluate_ValidExpressions(t *testing.T) {
	c := NewCalculator()

	tests := []struct {
		expression string
		expected   float64
	}{
		{"3 + 4", 7},
		{"10 + 2 * 6", 22},
		{"100 * 2 + 12", 212},
		{"100 * 2 + 12 / 4", 203},
		{"3 + 4 * 2", 11},
		{"18 / 3 - 2", 4},
		{"5", 5},
		{"3.5 + 2.5", 6},
	}

	for _, test := range tests {
		result, err := c.Evaluate(test.expression)
		if err != nil {
			t.Errorf("unexpected error for '%s': %v", test.expression, err)
			continue
		}
		if result != test.expected {
			t.Errorf("Evaluate(%q) = %v, expected %v", test.expression, result, test.expected)
		}
	}
}

func TestEvaluate_InvalidExpressions(t *testing.T) {
	c := NewCalculator()

	tests := []string{
		"",
		"     ",
		"3 +",
		"* 2",
		"1 /",
		"abc + 1",
	}

	for _, expr := range tests {
		_, err := c.Evaluate(expr)
		if err == nil {
			t.Errorf("expected error for invalid expression '%s', but got none", expr)
		}
	}
}
