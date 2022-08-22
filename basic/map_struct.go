package main

import (
	"fmt"
	"reflect"
)

// struct
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// embedding 很像繼承但其實是composition

type Animal struct {
	Name   string `required max: "100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// map
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"CA": 1,
		"TS": 2,
		"FA": 3,
	}
	// array is a valid key type but slice is not
	m := map[[3]int]string{}
	statePopulations["GA"] = 5
	delete(statePopulations, "FA")
	pop, ok := statePopulations["KK"]
	fmt.Println(pop, ok)
	fmt.Println(statePopulations, len(statePopulations), m)

	// struct
	aDoctor := Doctor{
		number:    3,
		actorName: "Tony",
		companions: []string{
			"Jack",
			"Lisa",
		},
	}

	fmt.Println(aDoctor.companions[1])

	oneDoctor := struct{ name string }{name: "John Pertwee"}
	anotherDoctor := oneDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(oneDoctor)
	fmt.Println(anotherDoctor)

	// embedding
	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "TW"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b.Name)

	// reflect package
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
