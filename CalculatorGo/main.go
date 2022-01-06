package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./operatorsFunc"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome, enter math operation (example 2 + 2) : \n ")
	scanner.Scan()
	operation := scanner.Text()

	result := 0

	if strings.Contains(operation, "+") {
		result = operatorsFunc.Sum(operation)
	} else if strings.Contains(operation, "-") {
		result = operatorsFunc.Subtract(operation)
	} else if strings.Contains(operation, "*") {
		result = operatorsFunc.Multiply(operation)
	} else if strings.Contains(operation, "/") {
		result = operatorsFunc.Division(operation)
	} else {
		fmt.Println("Error, check what you entered")

	}

	fmt.Println(result)

}
