package main

import (
	"fmt"
	"strconv"
)

// Person person
type Person struct {
	Name string
	Age  int
}

// "Stringer" is a embedded interface, just only implement function "String() string"
func main() {
	p := Person{"xiaoming", 20}
	fmt.Print(p)
}

func (p Person) String() string {
	return "Name:" + p.Name + ", Age:" + strconv.Itoa(p.Age)
}
