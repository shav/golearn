package main

import "fmt"

type Person struct {
	name string
}

// globalVar := 1 // global variables initialization is not supported
var globalVar int = 1

func main() {
	// Объявление и инициализация переменных
	//var num0 // error

	var num0 int
	fmt.Println(num0)

	var num1 int = 10
	fmt.Println(num1)

	var num2 = 20
	fmt.Println(num2)
	num2 = 22
	fmt.Println(num2)

	num3 := 30
	fmt.Println(num3)

	var (
		num4 = 40
		num5 = 50
	)
	fmt.Println(num4)
	fmt.Println(num5)

	num6, num7 := 60, 70
	fmt.Println(num6)
	fmt.Println(num7)

	// область видимости переменных
	fmt.Printf("globalVar: %d\n", globalVar)
	if true {
		var blockVar = 2
		fmt.Printf("blockVar if: %d\n", blockVar)
		for i := 0; i < 2; i++ {
			forVar := i
			fmt.Printf("blockVar for: %d\n", forVar)
		}
		// forVar = 3 // переменная из блока for недоступна вне блока
	} else {
		// blockVar = 3 // переменная из блока if недоступна автоматически в блоке else
		var blockVar = 3
		fmt.Printf("blockVar else: %d\n", blockVar)
	}
	// blockVar = 3 // переменная из блока if недоступна вне блока

	var f = func() {
		var funcVar = 4
		fmt.Printf("funcVar: %d\n", funcVar)
	}
	f()
	// funcVar = 5 // переменная из анонимной функции недоступна вне этой функции

	// Переменные и указатели на структуры
	var person0 Person
	fmt.Printf("%+v\n", person0)

	var personNil0 *Person
	fmt.Printf("%+v\n", personNil0)

	var personNil1 *Person = nil
	fmt.Printf("%+v\n", personNil1)

	var person1 = new(Person)
	fmt.Printf("%+v\n", person1)

	var person2 = Person{name: "Artem"}
	fmt.Printf("%+v\n", person2)

	var person3 *Person = &Person{name: "Artem"}
	fmt.Printf("%+v\n", person3)
}
