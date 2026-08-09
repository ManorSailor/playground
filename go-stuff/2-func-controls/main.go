package main

import (
	"errors"
	"fmt"
)

func main() {
	var answer, remainder, err = divide(10, 5)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Result is %v with remainder %v\n", answer, remainder)

	// normal switch case
	// breaks are optional
	// switch {
	// case answer%2 == 0:
	// 	fmt.Println("The answer is even")
	// default:
	// 	fmt.Println("IDK!")
	// }

	// Conditional switch
	// Different from other langs
	switch remainder {
	case 0:
		fmt.Println("Fully consumed")
	case 1, 2:
		fmt.Println("Looks like you're left with something.")
	default:
		fmt.Println("IDK!")
	}

	// Closures/First Class Funcs
	multiplyBy := multiplier(2)

	fmt.Println(multiplyBy(2))
	fmt.Println(multiplyBy(3))
	fmt.Println(multiplyBy(4))
}

func divide(num int, divisor int) (int, int, error) {
	var err error

	// Pretty much the same as JS. if, else if
	if divisor == 0 {
		err = errors.New("WTF are you trying to do? Go find a solution for division by 0.")
		return 0, 0, err
	}

	var answer int = num / divisor
	var remainder int = num % divisor

	return answer, remainder, err
}

// Functions are first class. Yay!
// Closures are there just like JS.
func multiplier(multiplyBy int) func(n int) int {
	return func(num int) int {
		return multiplyBy * num
	}
}
