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
	tomYaml, _ := yaml.Marshal(tom)
	fmt.Println(string(tomYaml))

	fmt.Println("-------------------------------------------------------------------")

	// Десериализация из json
	// Десериализация базовых типов
	numJson = []byte("2.71")
	json.Unmarshal(numJson, &num)
	fmt.Println(num)

	strJson = []byte("\"Hi,\\nArtem!\"")
	json.Unmarshal(strJson, &str)
	fmt.Println(str)

	// Десериализация в переменную несоответствующего типа
	strJson = []byte("Hello")
	var num1 int
	err := json.Unmarshal(strJson, &num1)
	fmt.Println(err)
	fmt.Println(num)

	fmt.Println("--------------------------------------")

	// Десериализация срезов
	sliceJson = []byte("[ 10,  11,12   ]")
	json.Unmarshal(sliceJson, &slice)
	fmt.Println(slice)

	fmt.Println("--------------------------------------")

	// Десериализация словарей
	mapJson = []byte(" {\"Artem\": 10,  \"Dimon\":   30 ,   \"Vovan\" : 20} ")
	json.Unmarshal(mapJson, &myMap)
	fmt.Println(myMap)

	fmt.Println("--------------------------------------")

	// Десериализация структур
	tomJson = []byte(`{
	  "Name": "Tom",
	  "Age": 20,
	    "Grades": [10,20,30 ],
	  "Contact":     {
		"Email": "tom@gmail.com",
		"Phone": "+7(965)123-456-78"
	  },
	  "Manager": {
		"Name": "Artem",
		"Age": 10,
		"Grades": null,
		"Contact": {
		  "Email": "",
		  "Phone": ""
		},
		"Manager": null
	  }
	}
	`)
	json.Unmarshal(tomJson, &tom)
	fmt.Println(tom)
	fmt.Println(tom.Manager)

	fmt.Println("--------------------------------------")

	// Десериализация из yaml
	tomYaml = []byte(`name: Tom
age: 24
grades:
    - 111
    - 222
    - 333
    - 777
contact:
    email: tom@gmail.com
    phone: "+1234567899"
manager:
    name: Artem
    age: 50
    grades: []
    contact:
        email: ""
        phone: ""
`)
	yaml.Unmarshal(tomYaml, &tom)
	fmt.Println(tom)
	fmt.Println(tom.Manager)
}
