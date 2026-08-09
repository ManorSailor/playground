package main

import (
	"fmt"
	"sync"
)

// This acts like a TS Branded Type, i.e., MyInt is a type of its own.
type MyInt uint8

// You can define custom methods on your own types.
func (myInt MyInt) square() MyInt {
	return myInt * myInt
}

// This is a type alias like TS, i.e., they can be used interchangeably.
// Only limited to a single type, no unions allowed.
type WaitGroup = sync.WaitGroup

// Generic Type Aliases are also possible since >=1.24. Nice.
type MyMap[K Number | string, V any] = map[K]V

// This defines a "group" or alias type for all possible Number type.
// Note: the ~ makes it an unbranded type, i.e., any underlying type can substitute for Number.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Generic Structs are also possible
type Base[T int | string | MyInt] struct {
	value T
}

// However, methods must not be generic themselves.
// Go will add this in newer version some time later
// func (b *Base[T]) inferVal[T int]() T {

func (b Base[T]) inferVal() {
	// This is a type switch in Go. It even works with custom types. Pretty Neat.
	// Cast to any is required... .type only works with interface
	switch any(b.value).(type) {
	case int:
		fmt.Printf("Integer here: %v\n", b.value)
	case string:
		fmt.Printf("String here: %v\n", b.value)
	case MyInt:
		fmt.Printf("MyInt here: %v\n", b.value)
	default:
		fmt.Printf("Unknown Shit.")
	}
}

func main() {
	var myInt MyInt

	var defInt uint8 = 5

	// Not possible to assign anything else to MyInt without cast
	// Even if it's the same underlying type.
	myInt = MyInt(defInt)

	myInt = myInt.square()

	fmt.Printf("Value of myInt: %v\n", myInt)

	fmt.Printf("\n")

	var intSlice []int8 = []int8{1, 2, 3, 4, 5}
	var f32Slice []float32 = []float32{1.1, 2.2, 3.3, 4.4, 5.5}

	// You cannot use "Group/Unbranded" types as variable types...
	// var t []Number = []Number{1, 2, 3}

	fmt.Printf("Sum of %v is %v\n", intSlice, sum(intSlice...))
	fmt.Printf("Sum of %v is %v\n", f32Slice, sum(f32Slice...))

	fmt.Printf("\n")

	// Must pass the specific type to the generic type, inference doesn't work here.
	base := Base[MyInt]{
		value: myInt,
	}

	base.inferVal()
}

// Generic function. Generics in go are compile-time only, i.e., the compiler generates the code
// using something called GC Stenciling. The runtime has no idea about generics. They are a compile-time concept.
func sum[T Number](nums ...T) T {
	var sum T

	for _, n := range nums {
		sum += n
	}

	return sum
}
