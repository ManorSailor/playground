// Exercise: Equivalent Binary Trees
// Link: https://go.dev/tour/concurrency/8

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	if t == nil {
		return
	}

	nodes := []*tree.Tree{}
	current := t

	for current != nil || len(nodes) > 0 {

		for current != nil {
			nodes = append(nodes, current)
			current = current.Left
		}

		n := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]

		ch <- n.Value

		current = n.Right
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, isC1Open := <-c1
		v2, isC2Open := <-c2

		if !isC1Open && !isC2Open {
			return true
		}

		if isC1Open != isC2Open {
			return false
		}

		if v1 != v2 {
			return false
		}
	}
}

func main() {
	t := tree.New(1)
	ch := make(chan int, 10)

	fmt.Println(t)

	go Walk(t, ch)

	for v := range ch {
		fmt.Printf("Tree Value: %v\n", v)
	}

	fmt.Printf("Are trees same? %v\n", Same(tree.New(1), tree.New(1)))
}
