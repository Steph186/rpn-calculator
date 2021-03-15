package calc

import (
	"errors"
	"fmt"
	"strconv"
)

var priorities = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "^": 3}

// compute operation between 2 int given the string operation
func HasCurrentLowerOrEqualPriority(current, popped string) bool {
	return priorities[current] <= priorities[popped]
}

func ConvertInflixToPostfix(inflix_expression string) (string, error) {
	if inflix_expression == "" {
		return "", nil
	}

	var current string
	var postfix_expression string
	var operator_stack []string

	for i := 0; i < len(inflix_expression); i++ {
		current = string(inflix_expression[i])

		_, err := strconv.Atoi(current)
		if err == nil {
			// CASE 1: INTEGER - push in postfix
			postfix_expression = fmt.Sprintf("%s%s", postfix_expression, current)

		} else if current == ")" {
			// CASE 2: Closing parenthesis - pop from operator stack until "("
			foundOpening := false

			for len(operator_stack) > 0 {
				// Pop element
				n := len(operator_stack)
				popped := operator_stack[n-1]
				operator_stack = operator_stack[:n-1]

				if popped == "(" {
					// End of loop: find the matching parenthesis
					foundOpening = true
					break
				} else {
					// Push in postfix
					postfix_expression = fmt.Sprintf("%s%s", postfix_expression, popped)
				}
			}

			if !foundOpening {
				return "", errors.New("Unbalanced parentheses")
			}

		} else if current == "(" || len(operator_stack) == 0 {
			// CASE 3: Opening parenthesis or empty operator stack
			operator_stack = append(operator_stack, current)

		} else {
			// CASE 4: Process Operations * / + -
			for len(operator_stack) > 0 {

				n := len(operator_stack) - 1
				last_pushed := operator_stack[n]
				value := HasCurrentLowerOrEqualPriority(current, last_pushed)

				if last_pushed == "(" || !value {
					// Parenthesis or higher priority operation cannot be popped
					break

				} else {
					// if lower priority element: pop it out of the stack and push into postfix exp
					postfix_expression = fmt.Sprintf("%s%s", postfix_expression, last_pushed)
					operator_stack = operator_stack[:n]
				}
			}

			// Add current operation to stack
			operator_stack = append(operator_stack, current)
		}
	}

	for i := range operator_stack {

		// Push remaining elements in stack in postfix
		var remaining string = string(operator_stack[len(operator_stack)-i-1])
		if remaining == "(" {
			return "", errors.New("Unbalanced parentheses")
		}

		postfix_expression = fmt.Sprintf("%s%s", postfix_expression, remaining)
	}

	return postfix_expression, nil
}
