package calc

import "testing"

func TestConversionEmpty(t *testing.T) {
	result := ConvertInflixToPostfix("")

	if result != "" {
		t.Errorf("Empty conversion failed, expected \"\", got %s", result)
	} else {
		t.Logf("Empty conversion success")
	}
}

func TestConversionUnbalanced(t *testing.T) {
	result := ConvertInflixToPostfix("(")

	if result != "" {
		t.Errorf("Unbalanced conversion failed, expected \"\", got %s", result)
	} else {
		t.Logf("Unbalanced conversion success")
	}
}

func TestConversionSum(t *testing.T) {
	result := ConvertInflixToPostfix("A+B")

	if result != "AB+" {
		t.Errorf("Conversion failed, expected \"AB+\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionSumMultiple(t *testing.T) {
	result := ConvertInflixToPostfix("A+B+C+D")

	if result != "AB+C+D+" {
		t.Errorf("Conversion failed, expected \"AB+C+D+\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionDivision(t *testing.T) {
	result := ConvertInflixToPostfix("A/B")

	if result != "AB/" {
		t.Errorf("Conversion failed, expected \"AB/\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionParenthesisSimple(t *testing.T) {
	result := ConvertInflixToPostfix("(A)+B")

	if result != "AB+" {
		t.Errorf("Conversion failed, expected \"AB+\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionParenthesis(t *testing.T) {
	result := ConvertInflixToPostfix("A/(B+C)")

	if result != "ABC+/" {
		t.Errorf("Conversion failed, expected \"ABC+/\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionCpmplex(t *testing.T) {
	result := ConvertInflixToPostfix("A+B*C/(E-F)")

	if result != "ABC*EF-/+" {
		t.Errorf("Conversion failed, expected \"ABC*EF-/+\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}
