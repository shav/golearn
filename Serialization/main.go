package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
)

type Contact struct {
	Email string
	Phone string
}

type Person struct {
	Name    string
	Age     int
	Grades  []int
	Contact Contact
	Manager *Person
}

func main() {
	// Сериализация в json
	// Сериализация базовых типов
	var num = 3.14
	numJson, _ := json.Marshal(num)
	fmt.Println(string(numJson))

	var str = "Hello,\nworld!"
	strJson, _ := json.Marshal(str)
	fmt.Println(string(strJson))

	fmt.Println("--------------------------------------")

	// Сериализация срезов
	var slice = []int{1, 2, 3, 4, 5}
	sliceJson, _ := json.Marshal(slice)
	fmt.Println(string(sliceJson))

	fmt.Println("--------------------------------------")

	// Сериализация словарей
	var myMap = map[string]int{"Artem": 1, "Vovan": 2, "Dimon": 3}
	mapJson, _ := json.Marshal(myMap)
	fmt.Println(string(mapJson))

	fmt.Println("--------------------------------------")

	// Сериализация структур
	var artem = Person{Name: "Artem", Age: 31}
	var vovan = Person{Name: "Vovan", Age: 29}
	var dimon = Person{Name: "Dimon", Age: 25}
	var tom = Person{
		Name: "Tom",
		Age:  24,
		Contact: Contact{
			Email: "tom@gmail.com",
			Phone: "+1234567899",
		},
		Grades:  []int{1, 2, 3},
		Manager: &artem,
	}
	tomJson, _ := json.MarshalIndent(tom, "", "  ")
	fmt.Println(string(tomJson))

	fmt.Println("--------------------------------------")

	// Сериализация среза структур
	var persons = []Person{tom, vovan, artem, dimon}
	personsJson, _ := json.MarshalIndent(persons, "", "  ")
	fmt.Println(string(personsJson))

	fmt.Println("--------------------------------------")

	// Сериализация в yaml
	tomJson, _ = yaml.Marshal(tom)
	fmt.Println(string(tomJson))
}
