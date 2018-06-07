package main

import (
	"fmt"
	"time"
)

// The basic goroutine, similar to create thread
func main() {
	// Using go + function to start a goroutine
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

// Show progress
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("%c", r)
			time.Sleep(delay)
		}
	}
}

// Calculate fib number
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
