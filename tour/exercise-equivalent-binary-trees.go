package main

import (
	"fmt"
	"reflect"
	"sort"

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

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	close(ch1)
	close(ch2)

	var a1 []int
	var a2 []int

	for i := range ch1 { 
		a1 = append(a1, i)
	}

	for i := range ch2 { 
		a2 = append(a2, i)
	}

	sort.Slice(a1, func(i, j int) bool {
		return a1[i] < a1[j]
	})
	sort.Slice(a2, func(i, j int) bool {
		return a2[i] < a2[j]
	})

	return reflect.DeepEqual(a1, a2)
}

func main() {
	// ch1 := make(chan int, 10)
	// ch2 := make(chan int, 10)

	t1 := tree.New(2)
	t2 := tree.New(1)

	fmt.Println(Same(t1, t2))



}
