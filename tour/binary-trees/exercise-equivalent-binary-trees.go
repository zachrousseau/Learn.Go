package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"

	"reflect"
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

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	var a1, a2 []int

	for i := range ch1 { 
		a1 = append(a1, i)
	}
	for i := range ch2 { 
		a1 = append(a2, i)
	}


	// Sort the slices
	sort.Ints(a1)
	sort.Ints(a2)

	fmt.Println(a1)
	fmt.Println(a2)

	// Compare the sorted slices
	return reflect.DeepEqual(a1, a2)
}

func main() {
	// ch1 := make(chan int, 10)
	// ch2 := make(chan int, 10)

	t1 := tree.New(1)
	t2 := tree.New(1)

	fmt.Println(Same(t1, t2))



}
