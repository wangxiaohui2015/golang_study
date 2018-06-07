package main

import (
	"fmt"
)

var java, python, c bool
var math, english, computer bool = true, false, true

// Using var to declare a variable
func main() {
	var i int
	var j = 1
	// Short variable declaration
	monday, tuesday, wednesday := 1, 2, 3
	fmt.Println(java, python, c, i)
	fmt.Println(j, math, english, computer)
	fmt.Println(monday, tuesday, wednesday)
}
