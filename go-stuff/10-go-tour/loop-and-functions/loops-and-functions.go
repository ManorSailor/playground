// Exercise: Loops and Functions
// Link: https://go.dev/tour/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	lastVal := z

	i := 0
	var isCloseEnough bool

	for !isCloseEnough {
		lastVal = z
		z -= (z*z - x) / (2 * z)

		fmt.Printf("Iteration: %v; Current value of z: %v\n", i+1, z)

		isCloseEnough = math.Abs(z-lastVal) < 0.001
		i++
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
}
