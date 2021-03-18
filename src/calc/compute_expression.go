package calc

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// compute operation between 2 int given the string operation
var operators = map[string]func(float64, float64) float64{
	"+": func(a, b float64) float64 { return a + b },
	"-": func(a, b float64) float64 { return a - b },
	"*": func(a, b float64) float64 { return a * b },
	"/": func(a, b float64) float64 { return a / b },
	"^": func(a, b float64) float64 { return math.Pow(a, b) },
}

func Calculator(rpx_expression string) (float64, error) {
	if rpx_expression == "" {
		return 0, nil
	}

	var rpx_array []string
	var current string
	var float_stack []float64
	var x1, x2, result float64

	rpx_expression = strings.Replace(rpx_expression, "(", "", -1)
	rpx_expression = strings.Replace(rpx_expression, ")", "", -1)
	rpx_array = strings.Split(rpx_expression, " ")

	for i := 0; i < len(rpx_array); i++ {
		current = string(rpx_array[i])

		c_int, err := strconv.Atoi(current)
		if err == nil {
			// Push integer in float stack
			float_stack = append(float_stack, float64(c_int))

		} else if strings.ContainsAny("+-/*^", current) {
			n := len(float_stack) - 1

			if n <= 0 && i != len(rpx_array) -1  {
				return 0, errors.New("Invalid expression")

			} else if n <= 0 && i == len(rpx_array) -1 {
				// Extra operation at the end of the expr
				return float_stack[0], nil
			}

			// Pop two elements
			x2, float_stack = float_stack[n], float_stack[:n]
			x1, float_stack = float_stack[n-1], float_stack[:n-1]

			// perform the operation and push result
			result = operators[current](x1, x2)
			float_stack = append(float_stack, float64(result))
		}
	}

	return float_stack[0], nil
}
