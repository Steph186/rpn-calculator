package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var priorities = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "^": 3}

// Compare priorities of current with previous ones
func HasCurrentLowerOrEqualPriority(current, popped string) bool {
	return priorities[current] <= priorities[popped]
}

// Check if next iteration is digit: if true, push in postfix and skip next it
func PushNextDigits(inflix string, postfix string, i int, lenght int) (string, int) {
	var err error
	for err == nil && i < lenght-1 {
		_, err = strconv.Atoi(string(inflix[i+1]))

		if err == nil {
			i = i+1
			postfix = fmt.Sprintf("%s%s", postfix, string(inflix[i]))
		}
	}

	return postfix, i
}

func ConvertInflixToPostfix(inflix string) (string, error) {
	if inflix == "" {
		return "", nil
	}

	var current string
	var postfix string
	var operator_stack []string
	var string_length int = len(inflix)

	// this variable keeps track of last char. If operation, means next "-" is unitary
	var is_unitary_operation bool = true

	for i := 0; i < string_length; i++ {
		current = string(inflix[i])

		if current == " " {
			continue
		}

		_, err := strconv.Atoi(current)
		if err == nil {
			// CASE 1: INTEGER - push in postfix
			postfix = fmt.Sprintf("%s %s", postfix, current)
			is_unitary_operation = false

			postfix, i = PushNextDigits(inflix, postfix, i, string_length)

		} else if current == ")" {
			// CASE 2: Closing parenthesis - pop from operator stack until "("
			foundOpening := false
			is_unitary_operation = false

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
					postfix = fmt.Sprintf("%s %s", postfix, popped)
				}
			}

			if !foundOpening {
				return "", errors.New("Unbalanced parentheses")
			}

		} else if current == "(" {
			// CASE 3: Opening parenthesis
			operator_stack = append(operator_stack, current)
			is_unitary_operation = true

		} else if current == "-" && is_unitary_operation {
			// CASE 4: unitary - . Transform to binary and skip next digit iteration
			postfix = fmt.Sprintf("%s 0 ", postfix) // add 0
			postfix, i = PushNextDigits(inflix, postfix, i, string_length)
			postfix = fmt.Sprintf("%s -", postfix) // add binary -
			is_unitary_operation = false

		} else {
			// CASE 5: Process Operations * / + -
			for len(operator_stack) > 0 {

				n := len(operator_stack) - 1
				last_pushed := operator_stack[n]
				value := HasCurrentLowerOrEqualPriority(current, last_pushed)

				if last_pushed == "(" || !value {
					// Parenthesis or higher priority operation cannot be popped
					break

				} else {
					// if lower priority element: pop it out of the stack and push into postfix exp
					postfix = fmt.Sprintf("%s %s", postfix, last_pushed)
					operator_stack = operator_stack[:n]
				}
			}

			// Add current operation to stack
			operator_stack = append(operator_stack, current)
			is_unitary_operation = true
		}
	}

	for i := range operator_stack {

		// Push remaining elements in stack in postfix
		var remaining string = string(operator_stack[len(operator_stack)-i-1])
		if remaining == "(" {
			return "", errors.New("Unbalanced parentheses")
		}

		postfix = fmt.Sprintf("%s %s", postfix, remaining)
	}

	return strings.TrimSpace(postfix), nil
}
