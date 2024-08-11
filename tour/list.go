package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Add(item T) *List[T]{
	return &List[T]{next: l, val: item}
}

func PrintAll[T any](l *List[T]){
	for true {
		if l == nil{
			break
		}
		fmt.Print(l.val, " ")
		l = l.next
	} 
}

func main() {
	var l *List[string]

	l = l.Add("hello")

	PrintAll(l)
}

