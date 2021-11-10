// Mervan Kaya palinda21 uppg 2.4
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0 // Inialize series with F0
	b := 1 // Initialize series with F1
	c := 0
	return func() int {
		c, a, b = a, b, a+b // Fn = Fn-1 + Fn-2 according to formula
		return c            // Return Fn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
