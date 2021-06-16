package operatorsfunc

import (
	"fmt"
	"strconv"
	"strings"
)

func Subtract(operation string) int {
	values := strings.Split(operation, "-")
	result := 0

	for i := 0; i < len(values); i++ {
		num, error := strconv.Atoi(values[i])

		if error != nil {
			fmt.Println(error)
			fmt.Println("Error, verifique que fue lo ingresado")
		} else {
			if result == 0 {
				result = num
			} else {
				result -= num
			}
		}
	}
	return result
}
