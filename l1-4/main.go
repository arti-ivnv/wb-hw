package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	// Context was mainly implemented in Go to control state processes
	// Creating context that will be canceled by Ctrl + C (SIGINT)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop() // cancelation function to release signal resources

	// Assume we have 10 running goroutines
	var waiter sync.WaitGroup

	for i := range 10 {

		waiter.Add(1)

		// We need to monitore a context state in terms of termination signal
		go func(int) {
			select {
			// As soon as we received a cancelation signal we need to clean up and send done chanel to the main chanel
			case <-ctx.Done():
				// stopping goroutin
				fmt.Println("Received a SIGINT... Stopping a go routine... Shutting down")
				// Clean up
				waiter.Done()
				return
			// Perorm some logic
			default:
				defer waiter.Done()
				fmt.Printf("Goroutine[%d]: do some stuff \n", i)
			}
		}(i)

		time.Sleep(2 * time.Second)
	}

	// We also want to send a close chanel to the main chanel
	<-ctx.Done()
	fmt.Printf("Main routine received a SIGINT... Shutting down \n")
	time.Sleep(2 * time.Second)

	waiter.Wait()

}
