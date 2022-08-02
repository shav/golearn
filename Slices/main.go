package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// срезы - как динамические списки
	// инициализация через литерал
	var users = []string{"Tom", "Alice", "Kate"}
	fmt.Printf("%v, len = %d, capacity = %d\n", users, len(users), cap(users))

	users = append(users, "Bob")
	fmt.Printf("%v, len = %d, capacity = %d\n", users, len(users), cap(users))

	fmt.Println(users[3])
	//fmt.Println(users[5]) // error: index out of range (len <= index <= capacity)
	//fmt.Println(users[10]) // error: index out of range (index > capacity)

	users = delete(users, 0)
	fmt.Printf("%v, len = %d, capacity = %d\n", users, len(users), cap(users))

	fmt.Println("--------------------------------------")

	// инициализация через make
	var numbers []int = make([]int, 3 /* длина */)
	numbers[0] = 1
	numbers[1] = 2
	fmt.Printf("%v, len = %d, capacity = %d\n", numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 4)
	fmt.Printf("%v, len = %d, capacity = %d\n", numbers, len(numbers), cap(numbers))

	for _, num := range numbers {
		fmt.Printf("%d, ", num)
	}
	fmt.Println()

	fmt.Println("--------------------------------------")

	// инициализация через make с указанием ёмкости
	numbers = make([]int, 0 /* длина */, 10 /* ёмкость */)
	numbers = append(numbers, 0, 1, 2, 3, 4, 5)
	fmt.Printf("%v, len = %d, capacity = %d\n", numbers, len(numbers), cap(numbers))

	fmt.Println("---------------------------------")

	// Проверка на равенство
	//var users2 = []string{"Tom", "Alice", "Kate"}
	//var users2_ = []string{"Tom", "Alice", "Kate"}
	// fmt.Printf("slice == slice (sliced are equal by value): %v\n", users2 == users2_) // error: срезы нельзя сравнивать между собой
	var notInitializedSlice []string
	fmt.Printf("notInitializedSlice == nil: %v\n", notInitializedSlice == nil)
	emptySlice := []string{}
	fmt.Printf("emptySlice == nil: %v\n", emptySlice == nil)

	fmt.Println("--------------------------------------")

	// срезы - как окно в массив
	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
		"",
	}
	terrestrial := planets[0:4] // или planets[:4]
	giants := planets[4:8]
	gasGiants := giants[0:2]   // или giants[:2]
	iceGiants := giants[2:4:4] // или giants[2:]
	fmt.Println(planets)
	fmt.Println(terrestrial)
	fmt.Println(giants)
	fmt.Println(gasGiants)
	fmt.Printf("%v, len = %d, capacity = %d\n", iceGiants, len(iceGiants), cap(iceGiants))

	iceGiants[0] = "Uran"
	fmt.Println("\nAfter slice modified:")
	fmt.Println(planets)
	fmt.Println(giants)

	planets[2] = "Earth"
	fmt.Println("\nAfter base array modified:")
	fmt.Println(planets)
	fmt.Println(terrestrial)

	iceGiants = append(iceGiants, "Плутон")
	iceGiants[1] = "Neptun"
	fmt.Println("\nAfter reallocated slice modified:")
	fmt.Println(planets)
	fmt.Printf("%v, len = %d, capacity = %d\n", iceGiants, len(iceGiants), cap(iceGiants))

	iceGiants = giants[2:4:5 /*ёмкость*/]
	fmt.Printf("\n%v, len = %d, capacity = %d\n", iceGiants, len(iceGiants), cap(iceGiants))
	iceGiants = append(iceGiants, "Плутон")
	iceGiants[1] = "Neptun"
	fmt.Println("After slice modified without reallocation:")
	fmt.Println(planets)
	fmt.Printf("%v, len = %d, capacity = %d\n", iceGiants, len(iceGiants), cap(iceGiants))

	fmt.Println("--------------------------------------")

	// срез строки
	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Println(tune)
	//tune[0] = 'd' // error: срез строки изменять нельзя
	neptune = "Poseidon"
	fmt.Println(tune)

	fmt.Println("--------------------------------------")

	// методы срезов
	fmt.Println(planets)
	sort.StringSlice(planets[:]).Sort() // сортировка массива "на месте"
	fmt.Println("Sorted array:")
	fmt.Println(planets)

	fmt.Println("--------------------------------------")

	// срез как аргумент функции с переменным числом параметров
	fmt.Println(join(numbers...))

	// передача срезов в качестве аргумента функции (по значению)
	fmt.Println("Original slice:")
	fmt.Println(users)
	setValue(users, 0, "Том")
	fmt.Println("After slice  modified at func:")
	fmt.Println(users)
}

func delete(slice []string, index uint) []string {
	return append(slice[:index], slice[index+1:]...)
}

func join(numbers ...int) string {
	numStr := make([]string, len(numbers))
	for i := range numbers {
		numStr[i] = fmt.Sprintf("%d", numbers[i])
	}
	return strings.Join(numStr, ", ")
}

func setValue(slice []string, index uint, value string) {
	slice[index] = value
}
