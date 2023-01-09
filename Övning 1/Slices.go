// Mervan Kaya palinda21 uppg. 2.2
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	big := make([][]uint8, dy) // Big list 2D list of length dy containing empty sublists
	for i := range big {       // Iterate through each row, omit the value
		small := make([]uint8, dx) // Create sublist of length dx to represent one row
		for j := range small {
			small[j] = uint8((j + i) / 2) // Iterate through each column of the row (omit the value), compute the value of the color
		}
		big[i] = small // Change the empty list to the sublist
	}
	return big
}

func main() {
	pic.Show(Pic)
}
