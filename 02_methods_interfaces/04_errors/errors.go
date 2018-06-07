package main

import (
	"fmt"
	"math"
	"time"
)

// MyError myerror
type MyError struct {
	When time.Time
	What string
}

// "error" is a embedded interface, just implement method "Error() string".
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func (e *MyError) Error() string {
	return "Error occurred, " + e.When.String() + ", Reason:" + e.What
}

// Sqrt sqrt
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &MyError{time.Now(), "x must >= 0"}
	}
	return math.Sqrt(x), nil
}
