package main

import (
	"fmt"
)

type Contact struct {
	email string
	phone string
}

type Person struct {
	name    string
	age     int
	contact Contact
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

	// Вложенные структуры
	var tom = Person{
		name: "Tom",
		age:  24,
		contact: Contact{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	fmt.Println(tom)
	fmt.Println(tom.contact.phone)

	tom.contact.phone = "+7(123)456-78-90"
	fmt.Println(tom)

	fmt.Println("--------------------------------------")

	// Структуры являются типом-значением
	tom2 := tom
	tom2.name = "Tom2"
	fmt.Println("Copied struct:")
	fmt.Println(tom2)
	fmt.Println("Original struct:")
	fmt.Println(tom)
	fmt.Println()

	tom.setName("TOM")
	fmt.Println("set name (by value):")
	fmt.Println(tom)
	tom.setNameByRef("TOM")
	fmt.Println("set name (by ref):")
	fmt.Println(tom)
}

func (person Person) setName(name string) {
	person.name = name
}

func (person *Person) setNameByRef(name string) {
	person.name = name
}
