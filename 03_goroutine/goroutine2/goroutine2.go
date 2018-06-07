package main

import (
	"fmt"
	"time"
)

// How to use chanel for communication between goroutines
// Chanel is like linked block queue in Java
func main() {
	ch := make(chan int)
	go reader(ch)
	// Send data to chanel
	for i := 0; i < 100; i++ {
		ch <- i
	}
	// Close chanel, this is not necessary.
	close(ch)
	time.Sleep(time.Second)
}

// Read data from chanel
func reader(ch chan int) {
	// Method 1
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Chanel is closed.")
	// Method 2
	/* for {
		data, ok := <-ch
		if ok {
			fmt.Println(data)
		} else {
			fmt.Println("Chanel is closed.")
			break
		}
	} */
}
