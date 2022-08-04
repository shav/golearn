package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	// Сортировка списков элементов базовых типов
	intList := []int{4, 2, 3, 1}
	sort.Ints(intList)
	fmt.Println(intList)

	strList := []string{"b", "a", "d", "c"}
	sort.Strings(strList)
	fmt.Println(strList)

	fmt.Println("--------------------------------------")

	// Сортировка через компаратор
	var family = []Person{
		{"Alice", 23},
		{"Eve", 20},
		{"David", 2},
		{"Bob", 25},
	}
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Name < family[j].Name
	})
	fmt.Println(family)

	fmt.Println("--------------------------------------")

	// Кастомная сортировка
	sort.Sort(ByAge(family))
	fmt.Println(family)

	fmt.Println("--------------------------------------")

	// Сортирровка словарей по ключу
	studentMarks := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}
	keys := make([]string, 0, len(studentMarks))
	for student := range studentMarks {
		keys = append(keys, student)
	}
	sort.Strings(keys)

	for _, student := range keys {
		fmt.Println(student, studentMarks[student])
	}
}
