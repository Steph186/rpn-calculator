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

func TestConversionInvalid(t *testing.T) {
	_, err := Calculator("2+")

	if err == nil {
		t.Errorf("Invalid expression test failed, expected error but did not get any.")
	} else {
		t.Logf("Invalid expression success")
	}
}

func TestCalculatorSum(t *testing.T) {
	result, _ := Calculator("34+")

	if result != 7 {
		t.Errorf("Sum \"34+\" failed, expected 7, got %v", result)
	} else {
		t.Logf("Division \"34+\" success")
	}
}

func TestCalculatorDivision(t *testing.T) {
	result, _ := Calculator("92/")

	if result != 4.5 {
		t.Errorf("Division \"92/\" failed, expected 4.5, got %v", result)
	} else {
		t.Logf("Division \"92/\" success")
	}
}

func TestCalculatorComplex(t *testing.T) {
	result, _ := Calculator("38579357*+/--**")

	if result != -42.31578947368422 {
		t.Errorf("Complex operation failed, expected 42 < value < 43, got %v", result)
	} else {
		t.Logf("Complex operation success")
	}
}

func TestCalculatorUnbalanced(t *testing.T) {
	result, _ := Calculator("4/-")

	if result != 0 {
		t.Errorf("Unbalanced operation failed, expected 0, got %v", result)
	} else {
		t.Logf("Unbalanced operation success")
	}
}
