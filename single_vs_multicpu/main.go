package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 0
var wg sync.WaitGroup

func increment() {
	defer wg.Done()
	for range 1000 {
		counter++
	}
}

func main() {
	runtime.GOMAXPROCS(2) // Try changing this to >1

	routineCount := 50

	for range routineCount {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println("Counter =", counter)
}
