package operatorsFunc

import (
	"fmt"
	"strconv"
	"strings"
)

func Sum(operation string) int {
	values := strings.Split(operation, "+")
	result := 0

	for i := 0; i < len(values); i++ {
		num, error := strconv.Atoi(values[i])

		if error != nil {
			fmt.Println(error)
			fmt.Println("Error, check what you entered")
		} else {
			result += num
		}
	}
	return result
}
