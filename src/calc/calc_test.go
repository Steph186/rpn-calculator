package calc

import (
	"testing"
)

func TestRPN(t *testing.T) {
	cases := []struct {
		input  string
		output float64
	}{
		{"3 4 +", 7},
		{"3 4 - 5 + -", 4},
		{"3 4 5 * -", -17},
		{"3 4 5 * -", -17},
		{"3 (4 5 *) -", -17},
		{"3 20 -", -17},
		{"3 4 - 5 *", -5},
		{"5 3 4 - *", -5},
		{"(3 4 -) 5 *", -5},
	}
	for i, c := range cases {
		out, err := Calculator(c.input)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		if out != c.output {
			t.Errorf("%d: expected %f, got %f", i, c.output, out)
		}
	}
}

func TestInflix(t *testing.T) {
	cases := []struct {
		input  string
		output float64
	}{
		{"3+4", 7},
		{"3-4+5", 4},
		{"4*5-3", 17},
		{"-3+4*5", 17},
		{"(4*5)-3", 17},
		{"-3+(4*5)", 17},
		{"20-3", 17},
		{"-3+20", 17},
		{"(3-4)*5", -5},
		{"5*(3-4)", -5},
		{"5*-1", -5},
		{"-1*5*", -5},
	}
	for i, c := range cases {
		rpn, err := ConvertInflixToPostfix(c.input)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		out, err := Calculator(rpn)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		if out != c.output {
			t.Errorf("%d: expected %f, got %f", i, c.output, out)
		}
	}
}
