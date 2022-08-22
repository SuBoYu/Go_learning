package main

import (
	"fmt"
	"strconv"
)

var I int = 42

var (
	user1 string = "Tony"
	user2 string = "Jack"
)

var (
	counter int = 0
)

func main() {
	fmt.Printf("%v, %T\n", I, I)
	var i int = 27
	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
}
