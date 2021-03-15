package main

import (
	"bufio"
	"calc"
	"fmt"
	"os"
	"strings"
)

func ReadFromInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}

func main() {
	fmt.Println("Enter expression:")
	infixString, err := ReadFromInput()

	if err != nil {
		fmt.Println("Error from input:", err.Error())
		return
	}

	postfix, conversion_error := calc.ConvertInflixToPostfix(infixString)
	if conversion_error != nil {
		fmt.Println("Error from conversion:", conversion_error.Error())
		return
	}

	fmt.Println("Converted into RPN notation:", postfix)
	fmt.Println("Result:", calc.Calculator(postfix))

	return
}
