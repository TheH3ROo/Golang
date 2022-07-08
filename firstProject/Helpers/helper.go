package helper

import (
	"fmt"
	"time"
)

// FunctionTimer receives a function name and prints the execution time.
// type: 'defer FunctionTimer("FunctionName")' into the first line of the function
// to measure the execution time
func FunctionTimer(name string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s took %v", name, time.Since(start))
	}
}
