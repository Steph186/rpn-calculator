package main

import (
	"calc"
	"fmt"
)

func main() {
	fmt.Printf("Result: %v", calc.ConvertInflixToPostfix("A+B"))
}
