package main

import (
	"fmt"
)

func main() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	s := "hello go"
	for k, v := range s {
		fmt.Println(k, string(v))
	}
}
