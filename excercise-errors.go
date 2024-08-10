// Exercise: Errors
// Copy your Sqrt function from the earlier exercise and modify it to return an error value.

// Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

// Create a new type

// type ErrNegativeSqrt float64
// and make it an error by giving it a

// func (e ErrNegativeSqrt) Error() string
// method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

// Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

// Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.


package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string { // All an error function should do is return the error. Nothing else. 
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) { // can return either a float64 or an error

	if x < 0 { 
		return 0, ErrNegativeSqrt(x)
	}

	var v float64 

	for z := float64(1); z < 1000; z++ {

		var i float64 = z - 1

		z -= (z*z - x) / (2 * z)

		if z/i == 1 {
			v = z
			break
		}
	}
	return v, nil 
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
