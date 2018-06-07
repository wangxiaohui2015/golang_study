package main

import (
	"fmt"
)

// Person person
type Person struct {
	name string
	age  int
}

// Method is a function with a specified receiver.
func main() {
	person := Person{
		"xiaoming", 20,
	}
	person.run(10)

	var num MyFloat
	num = -12.4
	fmt.Println(num.abs())
}

// Method is a function just with a receiver
// The method and receiver must defined in one file
func (p Person) run(steps int) {
	fmt.Println("Person", p.name, "age", p.age, "runs steps", steps)
}

// MyFloat myfloat
type MyFloat float64

func (num MyFloat) abs() float64 {
	if num < 0 {
		return float64(-num)
	}
	return float64(num)
}
