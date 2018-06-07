package main

import (
	"fmt"
)

// The length of a1 cannot be changed
var a1 [5]int
var s1 []int

// Test array and slice
// Array has a fixed length, but slice is not
func main() {
	testArray()
	testSlices()
	testCreateSlices()
	testDeleteElement()
	testFunc(1, 2, 3, 4)
	a := []int{2, 4, 6, 8, 10}
	testFunc(a...)
}

func testArray() {
	fmt.Printf("The type of a1 is %T\n", a1)
	a1[0] = 0
	a1[1] = 1
	a1[2] = 2
	a1[3] = 3
	a1[4] = 4
	fmt.Println(a1)
	fmt.Println("The lenght of a1 is", len(a1))
}

func testSlices() {
	s1 = a1[:3]
	fmt.Println(s1)
	// If a slices pointers to an array, change slice will change the array elements
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println("The length of s1:", len(s1))
	fmt.Println("The capacity of s1:", cap(s1))
}

func testCreateSlices() {
	var s2 []int
	fmt.Println(s2)
	fmt.Println("The length of s2:", len(s2))
	fmt.Println("The capacity of s2:", cap(s2))

	// We can change the length of slices
	s2 = append(s2, 1, 2, 3, 4, 5)
	fmt.Println("The length of s2:", len(s2))
	fmt.Println("The capacity of s2:", cap(s2))

	// Create a slices with length and capacity
	var s3 = make([]int, 1, 20)
	fmt.Println(s3)
	s3 = append(s3, 1)
	s3 = append(s3, 1)
	s3 = append(s3, 1)
	s3 = append(s3, 1)
	fmt.Println("The length of s3:", len(s3))
	fmt.Println("The capacity of s3:", cap(s3))

	// Create a slices with initial elements
	s4 := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(s4)
	fmt.Println("The length of s4:", len(s4))
	fmt.Println("The capacity of s4:", cap(s4))

	// Nill slices
	var s5 []int
	if s5 == nil {
		fmt.Println("s5 is nil.")
	}
}

func testDeleteElement() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	index := 3
	a = append(a[:index], a[index+1:]...)
	fmt.Println(a)
	a = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a[index+1:])
}

func testFunc(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}
