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
	result, _ := ConvertInflixToPostfix("(")

	if result != "" {
		t.Errorf("Unbalanced conversion failed, expected \"\", got %s", result)
	} else {
		t.Logf("Unbalanced conversion success")
	}
}

func TestConversionSum(t *testing.T) {
	result, _ := ConvertInflixToPostfix("4+2")

	if result != "42+" {
		t.Errorf("Conversion failed, expected \"42+\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionSumMultiple(t *testing.T) {
	result, _ := ConvertInflixToPostfix("4+5+1+6")

	if result != "45+1+6+" {
		t.Errorf("Conversion failed, expected \"45+1+6+\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionDivision(t *testing.T) {
	result, _ := ConvertInflixToPostfix("6/8")

	if result != "68/" {
		t.Errorf("Conversion failed, expected \"68/\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionParenthesisSimple(t *testing.T) {
	result, _ := ConvertInflixToPostfix("(2)+9")

	if result != "29+" {
		t.Errorf("Conversion failed, expected \"29+\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionParenthesis(t *testing.T) {
	result, _ := ConvertInflixToPostfix("2*(5+(7+2))")

	if result != "2572++*" {
		t.Errorf("Conversion failed, expected \"2572++*\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionComplex(t *testing.T) {
	result, _ := ConvertInflixToPostfix("3+4*0/(9-2)")

	if result != "340*92-/+" {
		t.Errorf("Conversion failed, expected \"340*92-/+\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}

func TestConversionSuperComplex(t *testing.T) {
	result, _ := ConvertInflixToPostfix("(1+(2+3*2))*4+2")

	if result != "1232*++4*2+" {
		t.Errorf("Conversion failed, expected \"1232*++4*2+\", got %s", result)
	} else {
		t.Logf("Conversion success")
	}
}
