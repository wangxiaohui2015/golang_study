package main

import (
	"fmt"
)

// Vertex vertex
type Vertex struct {
	X, Y float64
}

// There are two reasons to use pointer receiver instead of variable,
// 1, pointer can modify the value that its receiver points to.
// 2, avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
// a.B or a.B(), here the a can be pointer or variable.
func main() {
	v := Vertex{3, 4}
	v.scale(10)
	fmt.Println(v.X, v.Y)
	v.scale2(10)
	fmt.Println(v.X, v.Y)
}

// This is value pass, will copy the original value, and not change it
func (v Vertex) scale(num float64) {
	v.X = v.X * num
	v.Y = v.Y * num
}

// This is pointer pass, will change the original value
// Recommend way, we need to pass pointer here.
func (v *Vertex) scale2(num float64) {
	v.X = v.X * num
	v.Y = v.Y * num
}
