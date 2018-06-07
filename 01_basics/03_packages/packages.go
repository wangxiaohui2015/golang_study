package main

import (
	"fmt"
	"math/rand"
)

// Define a package name at the top of program file, all files in the same folder should have a same package name.
func main() {
	fmt.Println("My favorite number is:", rand.Intn(10))
}
