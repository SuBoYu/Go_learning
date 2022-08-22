package main

import (
	"fmt"
)

func main() {
	// array
	a := [...]int{97, 85, 93}
	var students [3]string
	var identityMatrix [3][3]int
	students[0] = "Lisa"
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	b := a
	b[1] = 5
	c := &a
	c[0] = 3

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("identity Matrix: %v\n", identityMatrix)

	// slice
	k := []int{1, 2, 3, 4}

	v := k
	v[2] = 5

	fmt.Printf("k: %v\n", k)
	fmt.Printf("len: %v\n", len(k))
	fmt.Printf("cap: %v\n", cap(k))

	d := make([]int, 3, 100)
	fmt.Println(d)
	fmt.Printf("len: %v\n", len(d))
	fmt.Printf("cap: %v\n", cap(d))

	e := []int{}
	fmt.Println(e)
	fmt.Printf("len: %v\n", len(e))
	fmt.Printf("cap: %v\n", cap(e))
	e = append(e, 1)
	fmt.Println(e)
	fmt.Printf("len: %v\n", len(e))
	fmt.Printf("cap: %v\n", cap(e))
	e = append(e, 2, 3, 4, 5)
	fmt.Println(e)
	fmt.Printf("len: %v\n", len(e))
	fmt.Printf("cap: %v\n", cap(e))
	e = append(e, []int{2, 3, 4, 5}...)
	fmt.Println(e)
	fmt.Printf("len: %v\n", len(e))
	fmt.Printf("cap: %v\n", cap(e))

}
