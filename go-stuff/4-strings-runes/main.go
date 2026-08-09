package main

import (
	"fmt"
	"strings"
)

func main() {
	myStr := "résumé😒!"

	// Go uses UTF-8 to represent strings
	// If you loop over the string mentioned above, i.e., with non-ASCII chars in them
	// You will actually get their utf code-point
	for idx, char := range myStr {
		fmt.Printf("myStr: %c; Code-point: %v; Idx: %v\n", char, char, idx)
	}

	fmt.Printf("\n")

	// It skips the index because some chars in the string require more than 1 byte
	// Hence, range skips over them to not give you corrupted chars.
	// Cast strings to []rune to avoid problem with index skips
	for idx, char := range []rune(myStr) {
		fmt.Printf("myStr: %c; Code-point: %v; Idx: %v\n", char, char, idx)
	}

	fmt.Printf("\n")

	// Two ways to build strings
	// Because they are immutable. Phew.
	myNewStr := ""

	for _, char := range myStr {
		myNewStr += strings.ToUpper(string(char))
	}

	fmt.Printf("myNewStr: %v\n", myNewStr)

	// Using StringBuilder, more verbose but efficient
	var myStrBuilder strings.Builder

	for _, char := range myStr {
		myStrBuilder.WriteRune(char)
	}

	myNewStr2 := myStrBuilder.String()

	fmt.Printf("myNewStr2: %v\n", myNewStr2)

	// Reversing a string. There is no in-built utility
	// Go's Maintainers: This doesn't seem necessary. Closing...
	// and yet, there are so many gotchas around reversing strings
	var myRevStrBuilder strings.Builder
	myNewStrRunes := []rune(myNewStr)

	for i := len(myNewStrRunes) - 1; i >= 0; i-- {
		myRevStrBuilder.WriteRune(myNewStrRunes[i])
	}

	fmt.Printf("Reversed myNewStr: %v\n", myRevStrBuilder.String())
}
