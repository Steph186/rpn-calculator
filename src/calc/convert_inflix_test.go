package calc

import "testing"

func TestIsLower1(t *testing.T) {
	result := HasCurrentLowerOrEqualPriority("-", "/")

	if result != true {
		t.Errorf("TestIsLower1 failed, expected true, got %v", result)
	} else {
		t.Logf("TestIsLower: higer popped, success")
	}
}

func TestIsLower3(t *testing.T) {
	result := HasCurrentLowerOrEqualPriority("/", "/")

	if result != true {
		t.Errorf("TestIsLower1 failed, expected true, got %v", result)
	} else {
		t.Logf("TestIsLower: equal, success")
	}
}

func TestConversionEmpty(t *testing.T) {
	result, _ := ConvertInflixToPostfix("")

	if result != "" {
		t.Errorf("Empty conversion failed, expected \"\", got %s", result)
	} else {
		t.Logf("Empty conversion success")
	}
}

func TestConversionUnbalanced(t *testing.T) {
	_, err := ConvertInflixToPostfix("(")

	if err == nil {
		t.Errorf("Unbalanced conversion failed, expected error but did not get any.")
	} else {
		t.Logf("Unbalanced conversion success")
	}
}

func TestConversionSum(t *testing.T) {
	result, _ := ConvertInflixToPostfix("4+2")

	if result != "4 2 +" {
		t.Errorf("Conversion failed, expected \"4 2 +\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionSumMultiple(t *testing.T) {
	result, _ := ConvertInflixToPostfix("4+ 5+1+ 6")

	if result != "4 5 + 1 + 6 +" {
		t.Errorf("Conversion failed, expected \"4 5 + 1 + 6 +\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionDivision(t *testing.T) {
	result, _ := ConvertInflixToPostfix("6 /8")

	if result != "6 8 /" {
		t.Errorf("Conversion failed, expected \"6 8 /\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionNumbers(t *testing.T) {
	result, _ := ConvertInflixToPostfix("238-80")

	if result != "238 80 -" {
		t.Errorf("Conversion failed, expected \"238 80 -\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionUnaryOperation(t *testing.T) {
	result, _ := ConvertInflixToPostfix("-80-(-1*4)")

	if result != "0 80 - 0 1 - 4 * -" {
		t.Errorf("Conversion failed, expected \"0 80 - 0 1 - 4 * -\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionParenthesisSimple(t *testing.T) {
	result, _ := ConvertInflixToPostfix("(2)+9")

	if result != "2 9 +" {
		t.Errorf("Conversion failed, expected \"2 9 +\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionParenthesis(t *testing.T) {
	result, _ := ConvertInflixToPostfix("2*(55+(7+2))")

	if result != "2 55 7 2 + + *" {
		t.Errorf("Conversion failed, expected \"2 55 7 2 + + *\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionComplex(t *testing.T) {
	result, _ := ConvertInflixToPostfix("3+4*0/(9-2)")

	if result != "3 4 0 * 9 2 - / +" {
		t.Errorf("Conversion failed, expected \"3 4 0 * 9 2 - / +\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionSuperComplex(t *testing.T) {
	result, _ := ConvertInflixToPostfix("(1+(2+3*2))*4^2+2")

	if result != "1 2 3 2 * + + 4 2 ^ * 2 +" {
		t.Errorf("Conversion failed, expected \"1 2 3 2 * + + 4 2 ^ * 2 +\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionUnitaryOp(t *testing.T) {
	result, _ := ConvertInflixToPostfix("5*-1")

	if result != "5 0 1 - *" {
		t.Errorf("Conversion failed, expected \"5 0 1 - *\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}
