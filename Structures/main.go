package main

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	Email string
	Phone string
}

type Person struct {
	Name    string
	Age     int
	Contact Contact
	Manager *Person
}

// Методы структур
func (person Person) setName(name string) {
	person.Name = name
}

func (person *Person) setNameByRef(name string) {
	person.Name = name
}

func main() {
	// Инициализация
	var person = Person{}
	fmt.Println(person)

	person = Person{Name: "Artem"}
	fmt.Println(person)

	person = Person{Name: "Artem", Age: 31}
	fmt.Println(person)

	fmt.Println("--------------------------------------")

	// Вложенные структуры
	var tom = Person{
		Name: "Tom",
		Age:  24,
		Contact: Contact{
			Email: "tom@gmail.com",
			Phone: "+1234567899",
		},
	}
	fmt.Println(tom)
	fmt.Println(tom.Contact.Phone)

	tom.Contact.Phone = "+7(123)456-78-90"
	fmt.Println(tom)

	fmt.Println("--------------------------------------")

	// Структуры являются типом-значением
	tom2 := tom
	tom2.Name = "Tom2"
	fmt.Println("Copied struct:")
	fmt.Println(tom2)
	fmt.Println("Original struct:")
	fmt.Println(tom)
	fmt.Println()

	tom.setName("TOM")
	fmt.Println("set Name (by value):")
	fmt.Println(tom)
	tom.setNameByRef("TOM")
	fmt.Println("set Name (by ref):")
	fmt.Println(tom)

	fmt.Println("--------------------------------------")

	// Вложенные ссылки на структуры
	tom.Manager = &person
	tomStr, _ := json.MarshalIndent(tom, "", "  ")
	fmt.Println(string(tomStr))

	fmt.Println("--------------------------------------")
}
