// Mervan Kaya palinda21 Task 1 Bug 2
package main

import (
	"fmt"
	"sync"
)

// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1) // Add one group to wait for
	go Print(ch)
	go func() {
		for i := 1; i <= 11; i++ {
			ch <- i // Writes as the same time as Print reads! Data Race!
		}
		wg.Done() // Decrements wg when finished with adding to channel
	}()
	wg.Wait() // Blocks until wg is decremented to 0
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
}
