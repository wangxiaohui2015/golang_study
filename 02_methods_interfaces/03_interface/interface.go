package main

import (
	"fmt"
)

// Runnable could run
type Runnable interface {
	Run(steps int)
}

// Person person
type Person struct {
	name string
	age  int
}

// Duck duck
type Duck struct {
	color string
	age   int
}

// The usage of interface, no "implemented" keyword in golang, to implement a interface, just only implement all functions in interface.
func main() {
	p := Person{"xiaoming", 10}
	run10steps(&p)
	d := Duck{"red", 2}
	run10steps(d)

	var r Runnable
	// Assign a variable to interface
	r = &p
	// Invoke method
	r.Run(20)
	r = d
	r.Run(30)

	// Interface{} can take data with any type, which is the same as Object in Java.
	var i interface{}
	i = 100
	fmt.Println(i)
	var num = i.(int)
	fmt.Println(num)
	i = "hello"
	fmt.Println(i)
	s, ok := i.(string)
	fmt.Println(s, ok)
}

// Run function is on type *Person, not the same as Person
func (p *Person) Run(steps int) {
	fmt.Println("Person", p.name, "runs steps", steps)
}

func (d Duck) Run(steps int) {
	fmt.Println("Duck, color is", d.color, "runs steps", steps)
}

func run10steps(r Runnable) {
	r.Run(10)
}
