package main

import (
	"RedBlackTree/RBTree"
	"fmt"
)

type set int

func (s set) Compare(se RBTree.Entryer) int {
	sh := se.(set)
	if s > sh {
		return -1
	} else if s < sh {
		return 1
	} else {
		return 0
	}
}

func (s set) GetValue() interface{} {
	return s
}

func main() {
	rb := RBTree.RBTree{}
	for i := 0; i < 20; i++ {
		s := set(i)
		rb.Insert(s)
	}
	rb.MidRec()
	fmt.Println()
	rb.LevelTraversal()
}
