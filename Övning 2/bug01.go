// Mervan Kaya palinda21 Task 1 Bug 1
package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string, 1) // Fix deadlock by buffering the channel
	ch <- "Hello world!"       // Blocked at send, need to receive in a seperate go routine or buffer the channel
	fmt.Println(<-ch)
}
