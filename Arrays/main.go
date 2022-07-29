package main

import (
	"fmt"
)

func main() {
	// инициализация - массив фиксированной длины
	var numbers5 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers5)

	var numbers52 [5]int = [5]int{1, 2}
	fmt.Println(numbers52)

	var numbers50 [5]int = [5]int{}
	fmt.Println(numbers50)

	var numbers2 [2]int = [2]int{1, 2}
	fmt.Println(numbers2)
	//numbers5 = numbers2 // error: длина массива является частью его типа

	fmt.Println("---------------------------------")

	// инициализация - массивы вычислимой длины
	var numbers5_ = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%v, len = %d\n", numbers5_, len(numbers5_))
	var numbers2_ = [...]int{1, 2, 3}
	fmt.Printf("%v, len = %d\n", numbers2_, len(numbers2_))

	// инициализация - массивы как словари
	colors := [3]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors)

	// инициализация - двумерный массив
	var table = [4][4]int{{1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}}
	fmt.Println(table)

	fmt.Println("---------------------------------")

	// обращение к элементам массива по индексу
	fmt.Printf("array[1]: %d\n", numbers5[1])
	numbers5[1] = 51
	fmt.Println(numbers5)
	// n7 := numbers5[7] // compile error: index out of range

	fmt.Println("---------------------------------")

	// заполнение - двумерный массив
	var board [4][4]int
	board[0][0] = 1
	board[0][3] = 1
	for column := range board[1] {
		board[1][column] = 1
	}
	fmt.Println(board)

	fmt.Println("---------------------------------")

	// итерация по элементам массива
	for i := len(numbers5) - 1; i >= 0; i-- {
		num := numbers5[i]
		fmt.Printf("%d, ", num)
	}
	fmt.Println()

	for _, num := range numbers5 {
		fmt.Printf("%d, ", num)
	}
	fmt.Println()

	fmt.Println("---------------------------------")

	// копирование массива
	copyNumbers5 := numbers5
	copyNumbers5[1] = 2
	fmt.Println("Copied changed array:")
	fmt.Println(copyNumbers5)
	fmt.Println("Original array:")
	fmt.Println(numbers5)

	fmt.Println("---------------------------------")

	// передача массива в качестве аргумента функции (по значению)
	fmt.Println("Original array:")
	fmt.Println(numbers5)
	setValue(numbers5, 0, -10)
	fmt.Println("After array  modified at func (by value):")
	fmt.Println(numbers5)
	// передача массива в качестве аргумента функции (по ссылке)
	setValueByRef(&numbers5, 0, -10)
	fmt.Println("After array  modified at func (by ref):")
	fmt.Println(numbers5)
}

func setValue(array [5]int, index uint, value int) {
	array[index] = value
}

func setValueByRef(array *[5]int, index uint, value int) {
	array[index] = value
}
