// Exercise: Slices
// Link: https://go.dev/tour/moretypes/18

package main

import (
	"math/rand/v2"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for i := range pic {
		picX := make([]uint8, dx)
		pic[i] = picX

		for i := range picX {
			picX[i] = uint8((i + 1) + (rand.IntN(dx)+1)/2)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
