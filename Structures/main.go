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
}
