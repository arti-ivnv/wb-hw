package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// Declaring and initializing input
	input := []int{2, 4, 6, 8, 10}
	// Creating a wait group to be able manage multiple threads
	// Some sort of semaphores as I got
	var waiter sync.WaitGroup

	// Iterating through the input array
	for _, value := range input {
		// Adding a new semaphore every interation
		waiter.Add(1)
		// Running an anonymous function in a separate thread aka goroutine
		go func(val int) {
			// Execute logic
			fmt.Println(int(math.Pow(float64(val), 2)))
			// Mark semapore as done as soon as logic will be proceeded
			waiter.Done()
		}(value)
	}

	// Waiting til all the go routines will be proceeded
	waiter.Wait()
}
