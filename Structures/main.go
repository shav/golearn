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

// Конструктор
func NewPerson(name string, age int) Person {
	return Person{name, age, Contact{}, nil}
}

// Методы структур
func (person Person) setName(name string) {
	person.Name = name
}

func (person *Person) setNameByRef(name string) {
	person.Name = name
}

// -------------------------------------------------------------------------------------

// Наследование
type Animal struct {
	Name string
}

func (animal *Animal) Eat() {
	fmt.Printf("%s eats\n", animal.Name)
}

func (animal *Animal) Sleep() {
	fmt.Printf("%s sleeps\n", animal.Name)
}

type Dog struct {
	Animal
}

func (dog *Dog) Bark() {
	fmt.Printf("%s barks\n", dog.Name)
}

type Cat struct {
	Animal
}

func (cat *Cat) Meow() {
	fmt.Printf("%s meows\n", cat.Name)
}

func (cat *Cat) Sleep() {
	cat.Animal.Sleep()
	fmt.Printf("Cat zzzzzz...")
}

func main() {
	// Инициализация
	var person = Person{}
	fmt.Println(person)

	person = Person{Name: "Artem"}
	fmt.Println(person)

	person = Person{Name: "Artem", Age: 31}
	fmt.Println(person)

	person = NewPerson("Artem", 20)
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

	// Проверка на равенство
	tom2 = tom
	fmt.Printf("struct == struct (equals): %v\n", tom == tom2)
	fmt.Printf("struct == struct (not equals): %v\n", tom == person)

	fmt.Println("--------------------------------------")

	// Вложенные ссылки на структуры
	tom.Manager = &person
	tomStr, _ := json.MarshalIndent(tom, "", "  ")
	fmt.Println(string(tomStr))

	fmt.Println("--------------------------------------")

	// Список из структур
	persons := []Person{person, tom}
	fmt.Println(persons)

	fmt.Println("--------------------------------------")

	// Наследование aka Композиция
	var cat = Cat{Animal{Name: "Murka"}}
	cat.Eat()
	cat.Meow()

	var dog = Dog{Animal{Name: "Rex"}}
	dog.Eat()
	dog.Bark()

	// var animal Animal = cat // error: наследование в go не настоящее :)
	// animal.Sleep()
}
