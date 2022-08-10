package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	// срезы - как окно в невидимый массив
	// инициализация через литерал
	var users = []string{"Tom", "Alice", "Kate"}
	fmt.Printf("%v, len = %d, capacity = %d\n", users, len(users), cap(users))

	// Добавление элемента в список
	users = append(users, "Bob")
	fmt.Printf("%v, len = %d, capacity = %d\n", users, len(users), cap(users))

	// Ловушка при добавлении в срез
	a := []byte("ba ")
	// HACK: добиваемся того, чтобы длина среза была меньше ёмкости (на некоторых машинах/ОС такое может быть изначально при создании среза)
	a[2] = 0
	a = a[:2] // В конце среза остаётся незаполненный элемент
	fmt.Printf("%v, len = %d, capacity = %d\n", a, len(a), cap(a))
	a1 := append(a, 'd')
	a2 := append(a, 'g')
	fmt.Println(string(a1)) // bag
	fmt.Println(string(a2)) // bag

	// Конкатенация списков
	var users1 = []string{"Artem", "Vovan", "Dimon"}
	users1 = append(users1, users...)
	fmt.Println(users1)

	fmt.Println(users[3])
	//fmt.Println(users[5]) // error: index out of range (len <= index <= capacity)
	//fmt.Println(users[10]) // error: index out of range (index > capacity)

	fmt.Println("Slice after delete item:")
	users = delete(users, 0)
	fmt.Printf("%v, len = %d, capacity = %d\n", users, len(users), cap(users))
	deleteInPlace(&users, 0)
	fmt.Printf("%v, len = %d, capacity = %d\n", users, len(users), cap(users))

	fmt.Println("--------------------------------------")

	// инициализация через make
	var numbers []int = make([]int, 3 /* длина */)
	numbers[0] = 1
	numbers[1] = 2
	fmt.Printf("%v, len = %d, capacity = %d\n", numbers, len(numbers), cap(numbers))

	// Добавление элементов в срез
	numbers = append(numbers, 4)
	fmt.Printf("%v, len = %d, capacity = %d\n", numbers, len(numbers), cap(numbers))

	// инициализация через make с указанием ёмкости
	numbers = make([]int, 0 /* длина */, 10 /* ёмкость */)
	numbers = append(numbers, 0, 10, 20, 30, 40, 50)
	fmt.Printf("%v, len = %d, capacity = %d\n", numbers, len(numbers), cap(numbers))

	fmt.Println("--------------------------------------")

	// Обход элементов среза в цикле
	for _, num := range numbers {
		fmt.Printf("%d, ", num)
	}
	fmt.Println()

	// Ловушка: обход индексов среза, а не значений!
	for i := range numbers {
		fmt.Printf("%d, ", i)
	}
	fmt.Println()

	// range вычисляется перед выполнением цикла -
	// все элементы, добавленные или удалённые во время итераций, в цикле не учитываются
	for index, num := range numbers {
		fmt.Printf("%d, ", num)
		numbers = append(numbers, 6+index)
	}
	fmt.Println()
	fmt.Println(numbers)

	for _, num := range numbers {
		fmt.Printf("%d, ", num)
		// Изменение циклической переменной никак не влияет элементы в самом срезе
		num = num * 2
	}
	fmt.Println()
	fmt.Println(numbers)

	fmt.Println("---------------------------------")

	// Копирование срезов
	fmt.Println("Copied slice:")
	var numbers0 []int
	copy(numbers0, numbers)
	fmt.Println(numbers0)

	var numbers2 = make([]int, 2)
	copy(numbers2, numbers)
	fmt.Println(numbers2)

	var numbers3 = make([]int, 20)
	copy(numbers3, numbers)
	fmt.Println(numbers3)

	fmt.Println("---------------------------------")

	// Проверка на равенство
	var users2 = []string{"Tom", "Alice", "Kate"}
	var users20 = []string{"Tom", "Alice", "Kate"}
	var users21 = []string{"Tom", "Alice", "Artem"}
	var users22 = []string{"Tom", "Alice"}
	// fmt.Printf("slice == slice (sliced are equal by value): %v\n", users2 == users20) // error: срезы нельзя сравнивать между собой через оператор ==
	fmt.Printf("slice == slice (equals): %v\n", Equal(users2, users20))
	fmt.Printf("slice == slice (not equals): %v\n", Equal(users2, users21))
	fmt.Printf("slice == slice (not equals): %v\n", Equal(users2, users22))

	var byteSlice1 = []byte("Hello")
	var byteSlice11 = []byte("Hello")
	var byteSlice2 = []byte("World")
	fmt.Printf("byte slice == byte slice (equals): %v\n", bytes.Equal(byteSlice1, byteSlice11))
	fmt.Printf("byte slice == byte slice (not equals): %v\n", bytes.Equal(byteSlice1, byteSlice2))

	var notInitializedSlice []string
	fmt.Printf("notInitializedSlice == nil: %v\n", notInitializedSlice == nil)

	emptySlice := []string{}
	fmt.Printf("emptySlice == nil: %v\n", emptySlice == nil)

	fmt.Println("--------------------------------------")

	// Линейный поиск элементов среза
	users = []string{"Tom", "Alice", "Kate"}
	fmt.Printf("existing: %t, item index: %d\n", Contains(users, "Kate"), Find(users, "Kate"))
	fmt.Printf("existing: %t, item index: %d\n", Contains(users, "Vovan"), Find(users, "Vovan"))

	// Бинарный поиск места для вставки элемента в отсортированный массив
	var sortedList = []int{1, 4, 6, 9, 10}
	fmt.Printf("existing item inserting index: %d\n", sort.SearchInts(sortedList, 4))
	fmt.Printf("non existing item inserting index: %d\n", sort.SearchInts(sortedList, 5))

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

	// передача срезов в качестве аргумента функции (по ссылке)
	fmt.Println("Original slice:")
	fmt.Println(users)
	setValue(users, 0, "Том")
	fmt.Println("After slice  modified at func:")
	fmt.Println(users)
}

func delete(slice []string, index uint) []string {
	return append(slice[:index], slice[index+1:]...)
}

func deleteInPlace(slice *[]string, index uint) {
	copy((*slice)[index:], (*slice)[index+1:]) // Shift slice[i+1:] left one index.
	(*slice)[len(*slice)-1] = ""               // Erase last element (write zero value).
	*slice = (*slice)[:len(*slice)-1]          // Truncate slice.
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

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Find(slice []string, searchValue string) int {
	for index, item := range slice {
		if searchValue == item {
			return index
		}
	}
	return -1
}

func Contains(slice []string, searchValue string) bool {
	for _, item := range slice {
		if searchValue == item {
			return true
		}
	}
	return false
}
