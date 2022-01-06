package operatorsFunc

import (
	"fmt"
	"strconv"
	"strings"
)

func Multiply(operation string) int {
	values := strings.Split(operation, "*")
	result := 0

	for i := 0; i < len(values); i++ {
		num, error := strconv.Atoi(values[i])

		if error != nil {
			fmt.Println(error)
			fmt.Println("Error, check what you entered")
		} else {
			if result == 0 {
				result = num
			}
			result *= num
		}
	}
	return result
}
