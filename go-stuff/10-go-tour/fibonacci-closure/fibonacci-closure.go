// Exercise: Fibonacci Closure
// Link: https://go.dev/tour/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	cur, next := 0, 1

	return func() int {
		result := cur
		cur, next = next, cur+next

		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
