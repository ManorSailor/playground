// Exercise: Error
// Link: https://go.dev/tour/methods/20

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.2f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	var lastVal float64
	var isCloseEnough bool

	for !isCloseEnough {
		lastVal = z
		z -= (z*z - x) / (2 * z)
		isCloseEnough = math.Abs(z-lastVal) < 0.001
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
