package main

import (
	"fmt"
	"log"
)

func main() {

	// defer
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	// panic
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}
