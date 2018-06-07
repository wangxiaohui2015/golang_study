package main

import (
	"fmt"
)

// The usage of function
func main() {
	fmt.Println("3 + 12 =", add(3, 12))
	fmt.Println("3 * 2 =", times(3, 2))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(17))
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

// Parameter share type
func times(num1, num2 int) int {
	return num1 * num2
}

// Multiple results
func swap(str1, str2 string) (string, string) {
	return str2, str1
}

// Named return value
func split(num int) (x, y int) {
	x = num * 4 / 9
	y = num - x
	return
}
