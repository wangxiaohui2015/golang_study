package main

import (
	"fmt"
	"sync"
	"time"
)

var num = 0

// Waiting all sub goroutines to finish.
func main() {
	mux := &sync.Mutex{}
	group := &sync.WaitGroup{}
	// Will start three goroutines, so add 3 in group
	group.Add(3)
	go add(mux, group, 10000000)
	go add(mux, group, 10000000)
	go add(mux, group, 10000000)
	// Wait util three goroutines are finished.
	group.Wait()
	fmt.Println(num)
}

func add(mux *sync.Mutex, group *sync.WaitGroup, n int) {
	// defer group done
	defer group.Done()
	mux.Lock()
	for i := 0; i < n; i++ {
		num++
	}
	mux.Unlock()
	time.Sleep(time.Second * 1)
	fmt.Println("Done.")
}
