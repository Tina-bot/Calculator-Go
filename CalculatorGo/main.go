package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Bienvenido, ingrese su operacion")
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
		fmt.Println("Error de operador")

	}

	fmt.Println(result)

}
