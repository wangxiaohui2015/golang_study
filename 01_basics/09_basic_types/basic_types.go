package main

import (
	"fmt"
	"math/cmplx"
)

var (
	flag = false
	// MaxInt The max int
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

// PI Constant variable
const PI = 3.14

// The basic types in golang
func main() {
	fmt.Println("flag=", flag)
	fmt.Println("The max uint64 is:", MaxInt)
	fmt.Println("z=", z)

	// Zero values
	var b bool
	var i int
	var s string
	var i32 int32
	var byt byte
	fmt.Println(b, i, s, i32, byt)

	// Convert value
	var intval = 10
	fmt.Println(intval)
	floadval := float32(intval)
	fmt.Println(floadval)

	// Variable inference
	num := 30.4
	fmt.Printf("The type of %f is:%T\n", num, num)

	fmt.Println("PI is", PI)
	fmt.Printf("The type of PI is:%T\n", PI)

	// Numeric constant
	const num2 = 1 << 100
	num3 := num2 >> 99
	fmt.Println(num3)
}
