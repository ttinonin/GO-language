package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

func createPerson(name string, age int) *Person {
	p := Person{name: name, age: age}

	return &p
}

func main() {
	daniel := Person{"Daniel", 20}
	pedro := createPerson("Pedro", 35)

	fmt.Println("Name:",daniel.name,"Age:",daniel.age)
	fmt.Println("Name:",pedro.name,"Age:",pedro.age)
}
