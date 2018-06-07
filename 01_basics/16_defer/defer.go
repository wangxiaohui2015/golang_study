package main

import (
	"fmt"
)

// The usage of defer
// defer is often used for close connection, unlock, etc.
func main() {
	fmt.Println("Begin main")
	testDefer()
	/*  After calling testDefer, all defer statements in testDefer function will be invoked one by one.
	 *	Like stack, the last defer statement will be invoked firstly. last-in-first-out
	 */
	fmt.Println("End main")
}

func testDefer() {
	fmt.Println("Begin testDefer")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("End testDefer")
}
