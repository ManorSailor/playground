package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Statically Typed
	// type is optional, normally inferrable
	var i int = 1

	// Strongly Typed as well. No TS needed.
	// i = "a"

	// Customizable data types...
	// There are many more than these... everything you except from a language
	var i16 int16 = 32_767                  // max int16
	var i32 int32 = 2_147_483_647           // max int32
	var f32 float32 = 16_777_216            // 2^24, precision limit for float32
	var f64 float64 = 9_007_199_254_740_992 // 2^53, precision limit for float64

	fmt.Printf("i = %v\n", i)
	fmt.Printf("i16 = %v\n", i16)
	fmt.Printf("i32 = %v\n", i32)
	fmt.Printf("f32 = %5f\n", f32)
	fmt.Printf("f64 = %5f\n", f64)

	// Arithmetic

	// different types require cast
	var f32_num float32 = 10.1
	var i32_num int32 = 2

	// intellisense auto-casts, but you can choose yourself too
	var f32_i32 = f32_num + float32(i32_num)
	var i32_f32 = i32_num + int32(f32_num)

	fmt.Printf("f32 + i32 = %3f\n", f32_i32)
	fmt.Printf("i32 + f32 = %v\n", i32_f32)

	// int division always rounds-off and gives int as result - same as any other lang
	// use % operator for remainder
	var i_num1 int = 10
	var i_num2 int = 2

	fmt.Printf("i_num1 / i_num2 = %v\n", i_num1/i_num2)
	fmt.Printf("i_num1 modulo i_num2 = %v\n", i_num1%i_num2)

	fmt.Println()

	// strings
	// can use escape sequences here as well
	// oh, and double quotes for strings, single for chars/runes
	var myStr string = "ABC!"
	// backticks for multi-lines, note: preserve spaces
	var myMultiLineStr string = `Hello
	World!`

	fmt.Println(myStr, myMultiLineStr)

	// len doesn't count length of a string, it returns bytes
	fmt.Printf("Bytes of a string via len(%s): %v\n", myStr, len(myStr))
	fmt.Printf("Bytes of a string via len(%s): %v\n", "😒", len("😒"))

	// use unicode/utf8 package for counting the length of chars
	fmt.Printf("Length of a string via utf8.RuneCountInString(%s): %v\n", "😒", utf8.RuneCountInString("😒"))

	// There is no char type. Go uses rune type for chars, consumes 4bytes to capture all unicode/utf-32 code points
	// note, it's just an alias for int32, you can replace rune with int32 without any issues. However, don't.
	var myRune rune = 'A'
	var myRune2 rune = '😒'

	fmt.Printf("MyRune1 (%c): %v\n", myRune, myRune)
	fmt.Printf("MyRune2 (%c): %v\n", myRune2, myRune2)

	// Just a tip, use walrus syntax (:=) to skip writing var,
	// but if you use that, you cannot define a data type
	num := 1

	// define vars in same line
	// can't define types tho. they are inferred.
	var1, var2 := 1, "b"

	fmt.Println(num, var1, var2)

	// constants
	// hello TS my old friend.
	// require initialization at definition.
	const PI float32 = 3.14

	fmt.Printf("PI: %v\n", PI)
}
