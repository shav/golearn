package main

import (
	"fmt"
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

	fmt.Println("--------------------------------------")

	// срез как аргумент функции с переменным числом параметров
	fmt.Println(join(numbers...))

	fmt.Println("--------------------------------------")
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
