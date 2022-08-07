package main

import (
	"fmt"
	"math"
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

	// Получение списка ключей и значений словаря
	keys := make([]string, 0, len(studentMarks))
	values := make([]int, 0, len(studentMarks))

	for k, v := range studentMarks {
		keys = append(keys, k)
		values = append(values, v)
	}
	fmt.Println(keys)
	fmt.Println(values)

	fmt.Println("---------------------------------")

	// так лучше не делать - получается полный треш, т.к. по словарям цикл итерируется в рандомном порядке
	// и по вновь добавленным элементам может как проитерироваться, так и нет
	for student, mark := range studentMarks {
		fmt.Printf("%s: %d\n", student, mark)
		studentMarks[student+"2"] = mark * 2
	}
	fmt.Println()
	fmt.Println(studentMarks)

	fmt.Println("---------------------------------")

	// lookup (по каждому ключу может храниться несколько значений):
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	tempGroups := make(map[int][]float64) // Словарь с ключами int и значениями []float64
	for _, t := range temperatures {
		g := int(math.Trunc(t/10) * 10)
		tempGroups[g] = append(tempGroups[g], t)
	}
	fmt.Println(tempGroups)
}

func setValue(dict map[string]int, key string, value int) {
	dict[key] = value
}
