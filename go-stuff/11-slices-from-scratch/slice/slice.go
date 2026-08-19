// Here we are trying to recreate slice from scratch
// Note: Due to internal language semantics, we cannot create internal semantics, i.e., [] being recognized as our custom slice
package slice

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

// Slice is internally a struct with the following structure
// Ref: https://go.dev/blog/slices-intro#Slice-internals
type Slice[T any] struct {
	Length, Capacity int
	// ptr *[]T -> This is a pointer to internal slice, not array.
	// unsafe.Pointer allows an arbitrary pointer to any type.
	// Go doesn't allow converting a different datatype pointer to another type pointer... unsafe.Pointer lets you
	// Eg: (*byte)(new(1)) is an error in Go; (*byte)(unsafe.Pointer(new(1))) is safe!
	ptr unsafe.Pointer
}

/**
* Use this to create Slice instances. Accepts optional length and capacity params using variadic params.
* Length must always the 1st element, and capacity should be right next to it.
 */
func Make[T any](args ...int) (Slice[T], error) {
	length := 0
	capacity := 1

	if len(args) == 1 {
		length = args[0]
		capacity = (length + 1)
	} else if len(args) == 2 {
		length = args[0]
		capacity = args[1]
	}

	if length < 0 || capacity <= 0 || capacity < length {
		return Slice[T]{}, errors.New("Please provide valid length and/or capacity for the Slice.")
	}

	// Go is Garbage Collected. Thus, we don't have malloc here - which also means that free'ing is not our headache.
	// This is the only way to get a block of memory of some type initialized. 
	// We then take a pointer to its 1st element because we will be manually moving around the allocated memory.
	mem := make([]T, capacity)
	ptr := unsafe.Pointer(&mem[0])

	return Slice[T]{
		Length:   length,
		Capacity: capacity,
		ptr:      ptr,
	}, nil
}

func (s Slice[T]) String() string {
	var str strings.Builder

	str.WriteRune('[')

	for idx := range s.Length {
		v, _ := s.At(idx)

		fmt.Fprint(&str, v)

		if idx != s.Length-1 {
			str.WriteString(", ")
		}
	}

	str.WriteRune(']')

	return str.String()
}

// append is a function which is used to insert values in a slice
// If there is capacity left in Slice.ptr, it appends to the END of the array
// Otherwise, it creates a new Array of (len(s.ptr) + len(v)), copies oldies then appends.
func (s *Slice[T]) Append(v ...T) {
	argsLen := len(v)
	availableRoom := s.Capacity - s.Length

	if argsLen > availableRoom {
		newSlice, _ := Make[T](s.Length, s.Capacity+argsLen)
		offset := 0

		for offset < s.Length {
			mem := newSlice.memAt(offset)

			*mem, _ = s.At(offset)

			offset++
		}

		s.ptr = newSlice.ptr
		s.Length = newSlice.Length
		s.Capacity = newSlice.Capacity
	}

	offset := s.Length

	for _, n := range v {
		mem := s.memAt(offset)
		*mem = n
		offset++
	}

	s.Length = offset
}

func (s *Slice[T]) Slice(bounds ...int) (Slice[T], error) {
	lower := 0
	upper := s.Length

	if len(bounds) == 1 {
		lower = bounds[0]
	} else if len(bounds) == 2 {
		lower = bounds[0]
		upper = bounds[1]
	}

	if lower < 0 || upper < lower || upper > s.Capacity {
		return Slice[T]{}, errors.New("Invalid lower/upper bounds, cannot slice.")
	}

	// Get the pointer at lower because that will be our first or starting index for this slice
	ptr := unsafe.Pointer(s.memAt(lower))

	return Slice[T]{
		Length:   upper - lower,
		Capacity: s.Capacity - lower,
		ptr:      ptr,
	}, nil
}

// Use At for index lookup of values.
// Supports negative indices as well.
func (s *Slice[T]) At(index int) (T, error) {
	var zero T

	if s.Length == 0 || index >= s.Length {
		return zero, errors.New("Index Out-Of-Bounds")
	}

	if index < 0 {
		index = ((index % s.Length) + s.Length) % s.Length
	}

	return *s.memAt(index), nil
}

// Insert at a value at an index - overwriting the existing value
func (s *Slice[T]) Insert(index int, val T) error {
	if index >= s.Length || index < 0 {
		return errors.New("Invalid index.")
	}

	mem := s.memAt(index)
	*mem = val

	return nil
}

func (s *Slice[T]) memAt(offset int) *T {
	var zero T
	size := unsafe.Sizeof(zero)

	return (*T)(unsafe.Add(s.ptr, uintptr(offset)*size))
}
