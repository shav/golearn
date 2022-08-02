package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	// Инициализация
	var person = Person{}
	fmt.Println(person)

	person = Person{name: "Artem"}
	fmt.Println(person)

	person = Person{name: "Artem", age: 31}
	fmt.Println(person)

	fmt.Println("--------------------------------------")
}
