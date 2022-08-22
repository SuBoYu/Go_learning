package main

import (
	"fmt"
)

func main() {
	// var n bool = true // true, false
	n := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	// float
	f := 3.14
	fmt.Printf("%v, %T\n", f, f)

	// complex
	var com complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", real(com), real(com))
	fmt.Printf("%v, %T\n", imag(com), imag(com))

	// string utf-8
	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", string(s[2]), s[2])
	fmt.Printf("%v, %T\n", b, b)

	// runes utf-32
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

}
