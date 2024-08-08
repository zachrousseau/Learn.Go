package main

import (
	"golang.org/x/tour/pic"

	"fmt"
)

func Pic(dx, dy int) [][]uint8 {

	result := make([][]uint8, dy) // Make a Matrix of len and cap of dy by dy


	fmt.Print("dy is ", dy)


	for y := 0; y < dy; y++{
		
		result[y] = make([]uint8, dx) // Make am array of length 5, cap 5 

		print(result[y])
	}

	print(result)

	return result 
}

func main() {
	pic.Show(Pic)
}
