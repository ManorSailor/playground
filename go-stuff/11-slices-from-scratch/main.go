package main

import (
	"fmt"
	"go-stuff/11-slices-from-scratch/slice"
)

func main() {
	fmt.Println("1. MAKE: length and capacity")

	builtin := make([]int, 2, 5)
	custom, _ := slice.Make[int](2, 5)

	fmt.Println("Built-in:", builtin)
	fmt.Println("Custom:  ", custom)

	fmt.Println("Built-in len:", len(builtin))
	fmt.Println("Built-in cap:", cap(builtin))

	fmt.Println("Custom len:  ", custom.Length)
	fmt.Println("Custom cap:  ", custom.Capacity)

	fmt.Println()

	// ------------------------------------------------------------
	fmt.Println("2. APPEND: add values")

	builtin = append([]int{}, 10, 20, 30)

	custom, _ = slice.Make[int](0, 3)
	custom.Append(builtin...)

	fmt.Println("Built-in:", builtin)
	fmt.Println("Custom:  ", custom)

	fmt.Println()

	// ------------------------------------------------------------
	fmt.Println("3. INDEX: read and overwrite")

	builtin = append([]int{}, 10, 20, 30)

	custom, _ = slice.Make[int](0, 3)
	custom.Append(builtin...)

	fmt.Println("Built-in index 1:", builtin[1])

	v, _ := custom.At(1)
	fmt.Println("Custom index 1:  ", v)

	builtin[1] = 99
	custom.Insert(1, 99)

	fmt.Println("Built-in:", builtin)
	fmt.Println("Custom:  ", custom)

	fmt.Println()

	// ------------------------------------------------------------
	fmt.Println("4. SLICE: create a view into the same backing array")

	builtin = append([]int{}, 10, 20, 30, 40, 50)

	custom, _ = slice.Make[int](0, 5)
	custom.Append(builtin...)

	// [low:high], where high is exclusive.
	builtinSliced := builtin[1:4]
	customSliced, _ := custom.Slice(1, 4)

	// Notice that the slices share different memory addresses, i.e., they are not the same
	fmt.Printf("Built-in original: %v; Len: %v; Cap: %v; Memory: %p\n", builtin, len(builtin), cap(builtin), &builtin)
	fmt.Printf("Built-in sliced: %v; Len: %v; Cap: %v; Memory: %p\n", builtinSliced, len(builtinSliced), cap(builtinSliced), &builtinSliced)
	fmt.Printf("Custom original: %v; Len: %v; Cap: %v; Memory: %p\n", custom, custom.Length, custom.Capacity, &custom)
	fmt.Printf("Custom sliced: %v; Len: %v; Cap: %v; Memory: %p\n", customSliced, customSliced.Length, customSliced.Capacity, &customSliced)

	// A slice is a view into an underlying array. Hence, This mutates the original.
	builtinSliced[0] = 99
	customSliced.Insert(0, 99)

	fmt.Println()

	// Slice expansion works too
	builtinSliced = builtin[:cap(builtin)]
	customSliced, _ = custom.Slice()

	// Notice, they are still DIFFERENT slices
	fmt.Printf("Built-in original: %v; Len: %v; Cap: %v; Memory: %p\n", builtin, len(builtin), cap(builtin), &builtin)
	fmt.Printf("Built-in sliced: %v; Len: %v; Cap: %v; Memory: %p\n", builtinSliced, len(builtinSliced), cap(builtinSliced), &builtinSliced)
	fmt.Printf("Custom original: %v; Len: %v; Cap: %v; Memory: %p\n", custom, custom.Length, custom.Capacity, &custom)
	fmt.Printf("Custom sliced: %v; Len: %v; Cap: %v; Memory: %p\n", customSliced, customSliced.Length, customSliced.Capacity, &customSliced)
}
