package main

import (
	"fmt"
	"tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := 0; i < 10; i++ {
		n1, n2 := <-c1, <-c2
		if n1 != n2 {
			return false
		}
	}
	return true
}

func main() {

	s1 := Same(tree.New(1), tree.New(1))
	fmt.Println(s1) // true

	s2 := Same(tree.New(1), tree.New(2))
	fmt.Println(s2) // false

}
