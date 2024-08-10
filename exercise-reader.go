// Exercise: Readers
// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
// "rewrite all values in []byte into 'A's"

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(b []byte) (n int, err error){

	for x := range b { 
		b[x] = 'A'
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
