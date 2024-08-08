// Exercise: Slices
// Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

// (You need to use a loop to allocate each []uint8 inside the [][]uint8.)

// (Use uint8(intValue) to convert between types.)



package main

import (
	"golang.org/x/tour/pic"
)


// dx = 5 
// dy = 256
func Pic(dx, dy int) [][]uint8 {

	result := make([][]uint8, dy) // Make a Matrix of len and cap of dy by dy

	for y := 0; y < dy; y++{
		
		result[y] = make([]uint8, dx) // Make am array of length 5, cap 5 


		for x:= 0; x < dx; x++{
			result[y][x] = uint8((x+y)) /2
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
