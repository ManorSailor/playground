package main

import "fmt"

func main() {
	// Real Arrays are back. GTFO JS & Python. Fixed. Contiguous. Same Type.

	// PS. I am aware of the ugly ass syntax. IDK who thought this reads better than:
	// int8 intArr[3] = {1, 2, 3}
	// or
	// int8 intArr[] = {1, 2, 3, 4}

	// Array of size 3 with default values of their respective data type
	var emptyIntArr [3]int8
	var emptyStrArr [3]string

	// Array of size 3 with values assigned
	var int8Arr [3]int8 = [3]int8{1, 2, 3}

	// Using the walrus guy saves you the duplication, but not the ugly syntax
	// int8Arr2 := [3]int8{1, 2, 3}

	// Array of inferred size from the values defined later
	// var int8Arr3 = [...]int8{4, 5, 6}
	int8Arr3 := [...]int8{4, 5, 6}

	fmt.Printf("emptyIntArr: %v\n", emptyIntArr)
	fmt.Printf("emptyStrArr: %v\n", emptyStrArr)

	fmt.Printf("int8Arr: %v\n", int8Arr)
	fmt.Printf("int8Arr3: %v\n", int8Arr3)

	// Fortunately, getting length is sane (unlike length of strings)
	fmt.Printf("Length of int8Arr3: %v\n", len(int8Arr3))

	fmt.Printf("\n")

	// Arrays are passed-by-value. Use pointers if you don't want a duplicate.
	// var copyOfInt8Arr = int8Arr
	copyOfInt8Arr := int8Arr

	// Mutating won't change the original
	copyOfInt8Arr[0] = 74

	fmt.Printf("Original int8Arr: %v at address: %p\n", int8Arr, &int8Arr)
	fmt.Printf("Copy of int8Arr: %v at address: %p\n", copyOfInt8Arr, &copyOfInt8Arr)

	fmt.Printf("\n")

	// What about vectors or lists? Well, Go has you covered with... Slice.
	// It's just a wrapper around array with a struct of Slice having 3 extra properties:
	// pointer to array, len & cap of that array.
	var emptyIntSlice []int
	var emptyStrSlice []string

	// Define a slice and assign values to it
	// var intSlice []int = []int{1, 2, 3}
	int8Slice := []int8{1, 2, 3}

	fmt.Printf("emptyIntSlice: %v\n", emptyIntSlice)
	fmt.Printf("emptyStrSlice: %v\n", emptyStrSlice)

	fmt.Printf("int8Slice: %v; Address: %p\n", int8Slice, &int8Slice)

	// There are 2 things: Length & Capacity of a slice.
	// Cap defines the total space available in the slice,
	// used to determine if copying is required before appending new element
	fmt.Printf("int8Slice Length: %v; Capacity: %v\n", len(int8Slice), cap(int8Slice))

	fmt.Printf("\n")

	// now we can append more elements into the slice
	// appending will either copy the array under the hood making more space,
	// or just append it if there is enough space
	int8Slice = append(int8Slice, 4)

	fmt.Printf("After appending int8Slice: %v; Address: %p\n", int8Slice, &int8Slice)
	fmt.Printf("int8Slice Length: %v; Capacity: %v\n", len(int8Slice), cap(int8Slice))

	// Introducing: Go's take on the spread operator
	// Like. Really. They have tried hard to avoid everything intuitive with other languages
	int8Slice2 := []int8{5, 6, 7, 8}
	mergedIntSlice := append(int8Slice, int8Slice2...)

	fmt.Printf("Merged int8Slice: %v; Address: %p\n", mergedIntSlice, &mergedIntSlice)

	// There is a helper to make your life easier
	var int8Slice3 []int8 = make([]int8, 3, 8)

	fmt.Printf("int8Slice3: %v; Length: %v; Capacity: %v\n", int8Slice3, len(int8Slice3), cap(int8Slice3))

	fmt.Printf("\n")

	// Onto Maps now.
	// I used to blame Python for having bad type annotations syntax. This clearly looks like a syntax error, but it's valid.
	var myMap map[string]int8 = make(map[string]int8)

	// Readability? What's that?
	myMap2 := map[string]int8{"Me": 26, "You": 40}

	fmt.Printf("myMap: %v; myMap2: %v\n", myMap, myMap2)

	// Accessing values is the same as any other language, similar to python
	fmt.Printf("My age: %v\n", myMap2["Me"])

	// Taking inspiration from JS's obj["non-existent"] == undefined, we have:
	fmt.Printf("nil's Age: %v\n", myMap2["nil"])

	// BUT. There is more: it's not limited to just ponzi JS returning undefined for everything
	// Each data type will return its flavor of "undefined", str will return "",
	// int will return 0 etc. Go doesn't have exceptions, remember?

	// How do we know if something exist or not? Surely, there must be .exists or something?
	// My boy. There is! The best part? Methods are for losers.
	var _, exists = myMap2["nil"]

	if !exists {
		fmt.Print("No shit? You don't want my undefined?!\n")
	}

	fmt.Printf("\n")

	// Loops.
	// Thou shalt only have for; for (no pun) thee have abused while, and do..while is cursed

	// What did go do here?
	// Some author at go: How can we improve the for loop?
	// THE LEAD: Um, I have always found the "()" to make my code too readable, let's get rid of just that.
	for name, age := range myMap2 {
		fmt.Printf("Name is %v & age is %v\n", name, age)
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("Yo %v!\n", i)
	}

	// Infinite Loop
	var i int8 = 0
	for {
		if i >= 3 {
			break
		}

		fmt.Printf("Yo %v from Infinity!\n", i)
		i += 1
	}
}
