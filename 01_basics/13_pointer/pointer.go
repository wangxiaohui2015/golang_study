package main

import (
	"fmt"
)

// The usage of pointer, pointer is just a memory address, and cannot operate pointer like C++.
func main() {
	var i, j = 10, 20
	// Declare a pointer with int format
	var p *int
	// Assign value to pointer
	p = &i
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(p)
	fmt.Println(*p)
	p = &j
	fmt.Println(*p)
	// Change value via pointer
	*p = 30
	fmt.Println(*p)
	fmt.Println(j)
}
