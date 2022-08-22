package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.123456789
	// 比較兩floating 大小
	if math.Abs(x/math.Pow(math.Sqrt(x), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are not the same")
	}

	switch 6 {
	case 1, 5, 10:
		fmt.Println("one")
	case 2, 4, 6:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

	var i interface{} = [3]int{}
	switch i.(type) {
	case int:
		fmt.Println("i is int")
	case float64:
		fmt.Println("i is float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}

}
