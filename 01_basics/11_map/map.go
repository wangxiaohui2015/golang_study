package main

import (
	"fmt"
)

// Person person
type Person struct {
	name string
	age  int
}

// The usage of map
func main() {
	var persons = map[string]Person{
		"1": Person{"xiaoming", 20},
		"2": Person{"tom", 30},
	}
	fmt.Println(persons["1"])
	fmt.Println(persons["2"])
	delete(persons, "1")
	fmt.Println(persons["1"])
	fmt.Println(persons["2"])
	p, ok := persons["1"]
	fmt.Println(p)
	// OK true: key "1" exists, false: not exists
	fmt.Println(ok)

	var myMap = make(map[string]int)
	myMap["1"] = 1
	myMap["2"] = 2
	myMap["3"] = 3
	fmt.Println(myMap["1"])
	fmt.Println(myMap["2"])
	fmt.Println(myMap["3"])

	// Delete a key-value
	delete(myMap, "2")
	num, ok := myMap["2"]
	fmt.Println(num, ok)
}
