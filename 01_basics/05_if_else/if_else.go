package main

import (
	"fmt"
)

func calc(num1, num2, num3 int) int {
	result := 0
	if flag := (num1 == 3); num3 == 1 {
		fmt.Println("Flag:", flag)
		result = num1 + num2
	} else if num3 == 2 {
		result = num1 - num2
	} else if num3 == 3 {
		result = num1 * num2
	} else {
		result = num1 / num2
	}
	return result
}

// The basic usage of if - else if - else
func main() {
	fmt.Println("3 + 2=", calc(3, 2, 1))
	fmt.Println("3 - 2=", calc(3, 2, 2))
	fmt.Println("3 * 2=", calc(3, 2, 3))
	fmt.Println("3 / 2=", calc(3, 2, 4))
}
