package main

import (
	"fmt"
)

// Using range in for
func main() {
	// Use both index and value on slices
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i, v := range s1 {
		fmt.Printf("Index is %d, value is %d\n", i, v)
	}
	// Only use index on array
	a1 := [12]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i := range a1 {
		fmt.Printf("Index is %d\n", i)
	}
	// Only use value on array
	for _, v := range a1 {
		fmt.Printf("Value is %d\n", v)
	}
}
