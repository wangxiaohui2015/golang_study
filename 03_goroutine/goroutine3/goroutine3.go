package main

import (
	"fmt"
	"time"
)

// How to use select to receive data from more chanells
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go reader(ch, quit)
	// Send data to chanel
	for i := 0; i < 100; i++ {
		ch <- i
	}
	// Send close status to goroutine
	quit <- true
	time.Sleep(time.Second)
}

// Read data from chanel
func reader(ch chan int, quit chan bool) {
	for {
		// Randomly select a case which is ready
		select {
		case data := <-ch:
			fmt.Println(data)
		case q := <-quit:
			if q {
				fmt.Println("Received close status.")
			}
		}
	}
}
