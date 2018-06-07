package main

import (
	"fmt"
	"sync"
	"time"
)

var num = 0

// Synchronize goroutines
func main() {
	mux := &sync.Mutex{}
	go add(mux, 10000000)
	go add(mux, 10000000)
	time.Sleep(time.Second * 1)
	// num is always 20000000, because two goroutines are synchronized
	fmt.Println(num)
}

func add(mux *sync.Mutex, n int) {
	// Like synchronized in Java, lock this func
	mux.Lock()
	// Unlock mux
	defer mux.Unlock()
	for i := 0; i < n; i++ {
		num++
	}
	fmt.Println("Done.")
}
