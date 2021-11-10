// Mervan Kaya palinda21 uppg 2.1
package main

import (
	"fmt"
	"math"
)

// Calc sqrt(x) using Newton's Method
func Sqrt(x float64) float64 {
	z := 1.0

	var znew float64
	diff := 1.0

	for diff > 1e-8 { // While diff less than tol
		znew -= (z*z - x) / (2 * z) // z+1
		diff = math.Abs(znew - z)   // Calc diff
		z = znew                    // z = z+1
	}
	return znew
}

func main() {
	fmt.Println(Sqrt(2))
}
