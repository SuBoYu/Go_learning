package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)

	k := [3]int{1, 2, 3}
	i := &k[0]
	j := &k[1]
	fmt.Printf("%v %p %p\n", k, i, j)

	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	(*ms).foo = 42
	fmt.Println(*ms)
}
