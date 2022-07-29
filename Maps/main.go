package main

import (
	"fmt"
)

func main() {
	// инициализация словарей
	studentMarks := make(map[string]int)
	fmt.Printf("%v, len = %d\n", studentMarks, len(studentMarks))

	studentMarks = make(map[string]int, 10)
	fmt.Printf("%v, len = %d\n", studentMarks, len(studentMarks))

	studentMarks = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 5,
	}
	fmt.Printf("%v, len = %d\n", studentMarks, len(studentMarks))

	// обращение к элементам словаря по ключу
	if val, ok := studentMarks["Tom"]; ok {
		fmt.Printf("Tom mark: %d\n", val)
	}
	// Если ключа в словаре нет, то возвращается значение по-умолчанию:
	fmt.Printf("Artem mark: %d\n", studentMarks["Artem"])

	studentMarks["Tom"] = 10
	fmt.Println(studentMarks)

	fmt.Println("---------------------------------")

	// словари передаются по ссылке
	studentMarks2 := studentMarks
	studentMarks2["Tom"] = 5
	fmt.Println(studentMarks)

	setValue(studentMarks, "Tom", 7)
	fmt.Println(studentMarks)

	fmt.Println("---------------------------------")

	// добавление элементов
	studentMarks["Artem"] = 5
	fmt.Println(studentMarks)

	// удаление элементов
	delete(studentMarks, "Alice")
	fmt.Println(studentMarks)

	fmt.Println("---------------------------------")

	// итерация по словарю
	for student, mark := range studentMarks {
		fmt.Printf("%s: %d\n", student, mark)
	}

	fmt.Println("---------------------------------")
}

func setValue(dict map[string]int, key string, value int) {
	dict[key] = value
}
