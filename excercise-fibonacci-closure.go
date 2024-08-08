// Exercise: Fibonacci closure
// Let's have some fun with functions.

// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {

	num1 := 0 
	num2 := 0

	return func(x int) int {

		num1 = num2 
		num2 = x 

		return num1 + num2 
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
