package main

import (
	"fmt"
	"sort"
	"slices"

	"golang.org/x/tour/tree"

	// "reflect"
)

func Same(t1, t2 *tree.Tree) bool {
	
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start the Walk function in a goroutine
	go func() {
		Walk(t1, ch1)
		close(ch1) // Close the channel after all values have been sent
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2) // Close the channel after all values have been sent
	}()
	

	// Add values into the slice
	var s1, s2 []int
	for i := range ch1 { 
		s1 = append(s1, i)
	}
	for i := range ch2 { 
		s2 = append(s2, i)
	}

	sort.Ints(s1)
	sort.Ints(s2)


	return slices.Equal(s1, s2)
}

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

	fmt.Println(Same(tree.New(1), tree.New(2)))
}
