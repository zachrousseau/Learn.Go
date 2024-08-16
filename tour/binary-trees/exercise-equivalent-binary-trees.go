package main

import (
	"fmt"
	// "sort"

	"golang.org/x/tour/tree"

	// "reflect"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}

	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)

}


func main() {
	ch := make(chan int)

	// Start the Walk function in a goroutine
	go func() {
		Walk(tree.New(1), ch)
		close(ch) // Close the channel after all values have been sent
	}()

	// Range over the channel to receive values
	for node := range ch {
		fmt.Println(node)
	}
}
