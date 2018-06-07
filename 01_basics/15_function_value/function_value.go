package main

import (
	"fmt"
	"math"
)

// Function is a variable, we can pass, assign function
func main() {
	// Define a anonymous function an assign it to a variable
	myFun := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// Call function via variable
	fmt.Println(myFun(5, 12))

	fmt.Println(compute(myFun))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// Call this function will return a function
func compute(fn func(num1, num2 float64) float64) float64 {
	return fn(3, 4)
}

// Function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
