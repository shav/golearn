package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
	"time"
)

type Contact struct {
	Email string
	Phone string
}

type Person struct {
	Name     string
	Age      int
	IsYoung  bool
	Grades   []int
	Contact  Contact
	Birthday DateTime
	Manager  *Person
}

type DateTime struct {
	time.Time
}

var defaultDateTime = time.Time{}

const universalDateTimeFormat = "2006-01-02T15:04:05-07:00"

func (datetime DateTime) MarshalJSON() ([]byte, error) {
	if datetime.Time == defaultDateTime {
		return []byte("\"\""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", datetime.Format(universalDateTimeFormat))
	return []byte(stamp), nil
}

func (datetime *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		datetime.Time = time.Time{}
		return
	}
	datetime.Time, err = time.Parse(universalDateTimeFormat, s)
	return
}

func (datetime DateTime) String() string {
	return datetime.Format(universalDateTimeFormat)
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
		Name:     "Tom",
		Age:      24,
		Birthday: DateTime{time.Date(1991, 1, 1, 9, 55, 0, 0, time.Local)},
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
	var tom2 Person
	tomJson = []byte(`{
	  "Name": "Tom",
	  "Age": 20,
	    "Grades": [10,20,30 ],
	  "Contact":     {
		"Email": "tom@gmail.com",
		"Phone": "+7(965)123-456-78"
	  },
      "Birthday": "2000-01-06T23:10:00+04:00",
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
	json.Unmarshal(tomJson, &tom2)
	fmt.Println(tom2)
	fmt.Println(tom2.Manager)

	fmt.Println("--------------------------------------")

	// Десериализация из yaml
	var tom3 Person
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
	yaml.Unmarshal(tomYaml, &tom3)
	fmt.Println(tom3)
	fmt.Println(tom3.Manager)
}
