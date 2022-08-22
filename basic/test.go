package main

import (
	"fmt"
	"strconv"
)

func main() {
	k := "7D"
	i, _ := strconv.Atoi(k[:len(k)-1])
	fmt.Println(i)
}
