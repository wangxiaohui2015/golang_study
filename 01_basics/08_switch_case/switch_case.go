package main

import (
	"fmt"
)

// Using switch - case
func main() {
	value := getValue("A")
	fmt.Println(value)
	result := calc(15, 10)
	fmt.Println(result)
	result = calc(15, 30)
	fmt.Println(result)
	result = calc(15, 40)
	fmt.Println(result)
}

func getValue(key string) string {
	switch value := ""; key {
	case "A":
		value = "The value of A is:" + "a"
		fmt.Println(value)
		return "a"
		// No need to add break here
	case "B":
		return "b"
	case "C":
		return "c"
	default:
		return "others"
	}
}

func calc(num1, num2 int) int {
	// This construct is similar to if-else statement, no condition after switch statement
	switch {
	case num1+10 > num2:
		return 10
	case num1+20 > num2:
		return 20
	default:
		return 0
	}
}
