// Mervan Kaya palinda21 Task 4
package main

import "fmt"

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan<- int) {
	sum := 0
	for _, val := range a {
		sum += val //Calculate sum
	}
	res <- sum // Add subtotal to channel
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	sub1 := <-ch // Get subtotal
	sub2 := <-ch // Get subtotal
	fmt.Printf("Sum is: %d", sub1+sub2)
}
