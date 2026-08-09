package main

import "fmt"

func main() {
	// Everything is pass by value in Go. Including Slices & Map.
	// Slices & Map are reference type, i.e., they use pointers underneath to point to respective structures.
	// Copying creates a new slice/map, but updating it modifies the original slice's array or map.

	// Pointer of type int8. `nil` unless assigned via new or &var
	// new sets the undefined of respective type at the memory location given by new
	var pint *int8 = new(int8)
	var num int8

	fmt.Printf("Value of pint: %v; %p\n", *pint, pint)
	fmt.Printf("Value of num: %v; %p\n", num, &num)

	// Same as C. Sets the value at the location pointed by pint to 3.
	*pint = 3

	fmt.Printf("Value of pint: %v; %p\n", *pint, pint)

	// Sets the pint pointer to the address of num
	pint = &num

	// Updates pint to 10 which will update num as well because both point to the same location
	*pint = 10

	fmt.Printf("Value of pint: %v; %p\n", *pint, pint)
	fmt.Printf("Value of num: %v; %p\n", num, &num)

	fmt.Printf("\n")

	// Slices/Map are also passed by Value, but they carry an internal pointer to Array or Map DS.
	var slice []int8 = []int8{1, 2, 3}
	slice_copy := slice

	// This will actually change the value of slice above...
	// Slice is just a Header struct with Pointer to array primitive, len & cap of array
	// It also handles allocation
	slice_copy[1] = 8

	fmt.Printf("Value of slice: %v; Address: %p; Cap: %v\n", slice, &slice, cap(slice))
	fmt.Printf("Value of slice_copy: %v; Address: %p; Cap: %v\n", slice_copy, &slice_copy, cap(slice_copy))

	fmt.Printf("\n")

	// What happens if you try to append past the cap of a slice?
	// Simple. It will allocate a new array larger array & copy all values to it
	s := make([]int8, 2)

	s[0] = 1
	s[1] = 2

	// Using index syntax will cause IndexOutOfRange error as it doesn't create a new array.
	// s[2] = 2

	fmt.Printf("Value of s before append: %v; Len: %v; Cap: %v\n", s, len(s), cap(s))

	s = append(s, []int8{3, 4}...)

	fmt.Printf("Value of s after append: %v; Len: %v; Cap: %v\n", s, len(s), cap(s))

	fmt.Printf("\n")
}
