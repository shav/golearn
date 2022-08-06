package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Sum[T Number](numbers ...T) T {
	var result T
	for _, num := range numbers {
		result = result + num
	}
	return result
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type List[T any] struct {
}

func (lst *List[T]) Add(item T) {
	fmt.Printf("Add item to list: %v\n", item)
}
func (lst *List[T]) GetAll() []T {
	return []T{*new(T)}
}

type IMovable interface {
	Move()
}

type Cat struct {
	Name string
}

func (cat Cat) Move() {
	fmt.Printf("%s moves\n", cat.Name)
}

type IAlive interface {
	Eat()
	SetName(name string)
}

type IAnimal interface {
	IAlive
	IMovable
}

func main() {
	var s = Sum(1, 2, 3)
	fmt.Println(s)

	var s2 = Sum[float64](1, 2, 3)
	fmt.Println(s2)

	// var s3 = Sum("a", "b", "c") // error: тип string не подходит под ограничение generic-типа

	fmt.Println("---------------------------------")

	studentMarks := map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 5,
	}
	keys := MapKeys(studentMarks)
	// keys := MapKeys[string, int](studentMarks)
	fmt.Println(keys)

	fmt.Println("---------------------------------")

	var list = List[int]{}
	list.Add(1)
	list.Add(2)
	fmt.Println(list.GetAll())

	// С ковариантностью/контравариантностью беда:
	// var listOfFloat List[float64] = list
	// var listOfAnimals List[IAnimal] = List[IAlive]{}
	// var listOfMovables List[IMovable] = List[Cat]{}

	// error: тип string не подходит под ограничение generic-типа
	// list.Add("str")
	// var listItems []string = list.GetAll()
}
