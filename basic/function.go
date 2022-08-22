package main

import (
	"fmt"
)

func sayMessage(name *string) {
	*name = "jack"
	fmt.Println(*name)
}

func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// method
type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	name := "tony"
	sayMessage(&name)
	fmt.Println(name)

	s := sum(1, 2, 3, 4, 5)
	fmt.Println(s)

	d, err := divide(5.0, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	f := func() {
		msg := "hello go"
		fmt.Println(msg)
	}
	f()

	g := greeter{
		greeting: "Hello",
		name:     "go",
	}
	g.greet()
}
