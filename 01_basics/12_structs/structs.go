package main

import (
	"fmt"
)

// Vector A vector
type Vector struct {
	X, Y int
}

// The usage of struct
func main() {
	v1 := Vector{2, 4}
	fmt.Println(v1)
	v1.X = 3
	fmt.Println(v1)
	var p *Vector
	// We often using pointer to operate struct
	p = &v1
	fmt.Println(p)
	fmt.Println(*p)
	v2 := Vector{X: 12, Y: 34}
	fmt.Println(v2)
	v3 := Vector{Y: 12}
	fmt.Println(v3)
	v4 := Vector{}
	fmt.Println(v4)
}
