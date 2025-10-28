package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func main() {

	// Check for recieced command-line input
	if len(os.Args) > 1 {
		fmt.Printf("amount of workers: %v \n", os.Args[1])
	} else {
		fmt.Println("no workers provided.. please enter re-run program with a comand-line argument")
		return
	}

	// Converting command-line into an integer variable
	workers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Errorf("Error converting string to int: %v", err)
	}

	// Allocating a main channel.
	// We also need to specify a buffer size to make sure that deadlock won't occure.
	mainChannel := make(chan int, workers)

	// Declaring a semaphore lock
	var waiter sync.WaitGroup

	// Sending random integers to the main channel.
	// Since we want n workers perform it's fair to assume that we will have n messages within the main channel
	for range workers {
		// for every new thread with an anonymous function add new lock
		waiter.Add(1)
		// send a rand num to the main chanel in a new thread implemeting anonymous
		go func() {
			// Open lock as soon as logic performed
			defer waiter.Done()

			var randNum = rand.Intn(1000)
			fmt.Println("Rand num generated: ", randNum)

			// Send random number to the main channel
			mainChannel <- randNum
		}()
	}

	// Wait till all the locks will be unlocked
	waiter.Wait()

	// Receiving messages from the main channel by a separate thread
	// Same semaphore logic implemented
	for range workers {
		waiter.Add(1)
		go func() {
			defer waiter.Done()
			fmt.Println(<-mainChannel)
		}()
	}

	waiter.Wait()

	// cleaning up by closing the channel
	close(mainChannel)

}
