package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t != nil {
		ch <- t.Value
		Walk(t.Left, ch)
		Walk(t.Right, ch)
	}
	
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	ch1 := make(chan int)

	t1 := tree.New(1)
	go Walk(t1, ch1)

	fmt.Println(t1)

	for i := range ch1 { 
		fmt.Println(i)
	}

}
