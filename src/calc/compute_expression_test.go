package calc

import (
	"testing"
)

func TestCalculatorEmpty(t *testing.T) {
	result, _ := Calculator("")

	if result != 0 {
		t.Errorf("Empty calculation failed, expected 0, got %v", result)
	} else {
		t.Logf("Empty division success")
	}
}

func TestCalculatorExtraLastOp(t *testing.T) {
	result, _ := Calculator("2 +")

	if result != 2 {
		t.Errorf("Sum \"3 4 +\" failed, expected 7, got %v", result)
	} else {
		t.Logf("Extra last op success")
	}
}

func TestCalculatorInvalid(t *testing.T) {
	_, err := Calculator("2 + +")

	if err == nil {
		t.Errorf("Expected invalid expression, did not get it")
	} else {
		t.Logf("Calculation invalid success")
	}
}

func TestCalculatorSum(t *testing.T) {
	result, _ := Calculator("3 4 +")

	if result != 7 {
		t.Errorf("Sum \"3 4 +\" failed, expected 7, got %v", result)
	} else {
		t.Logf("Division \"34+\" success")
	}
}

func TestCalculatorDivision(t *testing.T) {
	result, _ := Calculator("9 2 /")

	if result != 4.5 {
		t.Errorf("Division \"92/\" failed, expected 4.5, got %v", result)
	} else {
		t.Logf("Division \"92/\" success")
	}
}

func TestCalculatorComplex(t *testing.T) {
	result, _ := Calculator("3 8 5 7 9 3 5 7 * + / - - * *")

	if result != -42.31578947368422 {
		t.Errorf("Complex operation failed, expected 42 < value < 43, got %v", result)
	} else {
		t.Logf("Complex operation success")
	}
}

func TestCalculatorUnbalanced(t *testing.T) {
	result, _ := Calculator("4 / -")

	if result != 0 {
		t.Errorf("Unbalanced operation failed, expected 0, got %v", result)
	} else {
		t.Logf("Unbalanced operation success")
	}
}
