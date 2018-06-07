package main

import (
	"fmt"
	"time"
)

// A sharing variable
var num = 0

// Potential issue may occur when visit sharing variable in different goroutines
func main() {
	go add(10000000)
	go add(10000000)
	time.Sleep(time.Second * 1)
	// The num is not 20000000, because two goroutines may change num at the same time.
	fmt.Println(num)
}

func add(n int) {
	for i := 0; i < n; i++ {
		num++
	}
	fmt.Println("Done.")
}
