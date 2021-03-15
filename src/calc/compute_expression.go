package calc

import (
	"math"
	"strconv"
)

// compute operation between 2 int given the string operation
var operators = map[string]func(float64, float64) float64{
	"+": func(a, b float64) float64 { return a + b },
	"-": func(a, b float64) float64 { return a - b },
	"*": func(a, b float64) float64 { return a * b },
	"/": func(a, b float64) float64 { return a / b },
	"^": func(a, b float64) float64 { return math.Pow(a, b) },
}

func Calculator(rpx_expression string) float64 {
	if rpx_expression == "" {
		return 0
	}

	var current string
	var float_stack []float64
	var x1, x2, result float64

	for i := 0; i < len(rpx_expression); i++ {
		current = string(rpx_expression[i])

		i, _ := strconv.Atoi(current)

		switch current {
		case "+",
			"-",
			"/",
			"*",
			"^":

			n := len(float_stack) - 1
			if n <= 0 {
				return 0 // Unbalanced expression, calculation is not possible
			}

			// Pop two elements
			x2, float_stack = float_stack[n], float_stack[:n]
			x1, float_stack = float_stack[n-1], float_stack[:n-1]

			// perform the operation and push result
			result = operators[current](x1, x2)
			float_stack = append(float_stack, float64(result))
		default:
			// Push integer in float stack
			float_stack = append(float_stack, float64(i))
		}
	}

	return float_stack[0]
}
