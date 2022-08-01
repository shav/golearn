package main

import (
	"fmt"
	"math"
	"strings"
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

	fmt.Println("---------------------------------")

	// Проверка на равенство
	//studentMarks_ := map[string]int{
	//	"Tom":   1,
	//	"Bob":   2,
	//	"Sam":   4,
	//	"Alice": 5,
	//}
	// fmt.Printf("map == map (maps are equal by value): %v\n", studentMarks == studentMarks_) // error: словари нельзя сравнивать между собой
	var notInitializedMap map[string]int
	fmt.Printf("notInitializedMap == nil: %v\n", notInitializedMap == nil)
	emptyMap := make(map[string]int)
	fmt.Printf("emptyMap == nil: %v\n", emptyMap == nil)

	fmt.Println("---------------------------------")

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

	// lookup
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	// Словарь с ключами int и значениями []float64
	tempGroups := make(map[int][]float64)
	for _, t := range temperatures {
		g := int(math.Trunc(t/10) * 10)
		tempGroups[g] = append(tempGroups[g], t)
	}
	fmt.Println(tempGroups)

	fmt.Println("---------------------------------")

	set := NewSet()
	set.add(1, 2, 3)
	set.add(1, 2)
	fmt.Println(set)

	set.remove(1)
	fmt.Println(set)

	searchValue := 2
	fmt.Printf("set contains %d: %v\n", searchValue, set.contains(searchValue))
	searchValue = 1
	fmt.Printf("set contains %d: %v\n", searchValue, set.contains(searchValue))
}

func setValue(dict map[string]int, key string, value int) {
	dict[key] = value
}

type Set struct {
	items map[int]bool
}

func NewSet(values ...int) *Set {
	set := &Set{items: make(map[int]bool)}
	set.add(values...)
	return set
}

func (set *Set) contains(value int) bool {
	exists, ok := set.items[value]
	return ok && exists
}

func (set *Set) add(values ...int) {
	for _, value := range values {
		set.items[value] = true
	}
}

func (set *Set) remove(values ...int) {
	for _, value := range values {
		delete(set.items, value)
	}
}

func (set *Set) String() string {
	numStr := make([]string, len(set.items))
	i := 0
	for item, _ := range set.items {
		numStr[i] = fmt.Sprintf("%d", item)
		i++
	}
	return fmt.Sprintf("{%s}", strings.Join(numStr, ", "))
}
