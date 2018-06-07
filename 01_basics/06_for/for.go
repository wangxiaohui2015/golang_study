package main

import (
	"fmt"
)

// The usage of for, in golang only for statement, no wile statement.
func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1 + 2 + ... + 100 =", sum)

	sum = 0
	// Note, the initialize and post statement could be empty, it is while in Java.
	for sum < 100 {
		sum++
	}
	fmt.Println("Sum:", sum)

	for i := 0; ; i++ {
		if i == 2 || i == 4 {
			fmt.Println("i is 2 or 4.")
			// Using continue in for
			continue
		}
		if i > 10 {
			fmt.Println("i is larger than 10.")
			// Using break in for
			break
		}
		fmt.Println(i)
	}
}
