package calculator

import (
	"testing"

	"github.com/mnogu/go-calculator"
)

func TestCalculate(t *testing.T) {
	val, err := calculator.Calculate("1 + 2")
	if err != nil {
		t.Fatal(err)
	}

	if val != 3 {
		t.Errorf("Expected 3, got %f", val)
	}

	val, err = calculator.Calculate("1 - 2")
	if err != nil {
		t.Fatal(err)
	}

	if val != -1 {
		t.Errorf("Expected -1, got %f", val)
	}

	val, err = calculator.Calculate("1 * 2")
	if err != nil {
		t.Fatal(err)
	}

	if val != 2 {
		t.Errorf("Expected 2, got %f", val)
	}

	val, err = calculator.Calculate("1 / 2")
	if err != nil {
		t.Fatal(err)
	}

	if val != 0.5 {
		t.Errorf("Expected 0.5, got %f", val)
	}

	val, err = calculator.Calculate("1 - 2 + 3 * 4 / 5")
	if err != nil {
		t.Fatal(err)
	}

	if val != 1.4 {
		t.Errorf("Expected 2.4, got %f", val)
	}
}
