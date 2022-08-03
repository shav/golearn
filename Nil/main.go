package main

import (
	"fmt"
)

type IAlive interface {
	Eat()
}

type Person struct {
	Name string
}

func (person Person) Eat() {
	fmt.Printf("%s eats", person.Name)
}

func main() {
	// Срезы
	var slice []int
	fmt.Printf("no-init-slice == nil: %v\n", slice == nil)

	slice = nilSlice()
	fmt.Printf("nil-init-slice == nil: %v\n", slice == nil)

	// fmt.Println(slice[0]) // error: index out of range (не NullReferenceException)
	fmt.Println(slice[:])          // возвращает пустой врез []
	fmt.Println(append(slice, 10)) // вообще успещно добавляет элемент, без NRE

	slice = []int{}
	fmt.Printf("empty-slice == nil: %v\n", slice == nil)

	fmt.Println("--------------------------------------")

	// Словари
	var myMap map[string]int
	fmt.Printf("no-init-map == nil: %v\n", myMap == nil)

	myMap = nilMap()
	fmt.Printf("nil-init-map == nil: %v\n", myMap == nil)

	fmt.Println(myMap["Hello"])
	// myMap["Hello"] = 10 // error: NRE
	delete(myMap, "Hello")

	myMap = map[string]int{}
	fmt.Printf("empty-map == nil: %v\n", myMap == nil)

	fmt.Println("--------------------------------------")

	// Указатели на базовые типы
	var basicPointer *int
	fmt.Printf("no-init-basic-pointer == nil: %v\n", basicPointer == nil)

	basicPointer = nilBasicPointer()
	fmt.Printf("nil-init-basic-pointer == nil: %v\n", basicPointer == nil)

	// fmt.Println(*basicPointer) // NRE
	// *basicPointer = 10 // NRE

	basicPointer = (*int)(nil)
	fmt.Printf("typed-nil-init-basic-pointer == nil: %v\n", basicPointer == nil)

	fmt.Println("--------------------------------------")

	// Указатели на структуры
	var structPointer *Person
	fmt.Printf("no-init-struct-pointer == nil: %v\n", structPointer == nil)

	structPointer = nilStructPointer()
	fmt.Printf("nil-init-struct-pointer == nil: %v\n", structPointer == nil)

	// fmt.Println(*structPointer) // NRE
	// *structPointer = Person{}   // NRE

	structPointer = (*Person)(nil)
	fmt.Printf("typed-nil-init-struct-pointer == nil: %v\n", structPointer == nil)

	fmt.Println("--------------------------------------")

	// Указатели на интерфейсы
	var interfacePointer *IAlive
	fmt.Printf("no-init-interface-pointer == nil: %v\n", interfacePointer == nil)
	interfacePointer = nilInterfacePointer()
	fmt.Printf("nil-init-interface-pointer == nil: %v\n", interfacePointer == nil)
	interfacePointer = (*IAlive)(nil)
	fmt.Printf("typed-nil-init-interface-pointer == nil: %v\n", interfacePointer == nil)

	fmt.Println("--------------------------------------")

	// Интерфейсы
	var iface IAlive
	fmt.Printf("no-init-interface == nil: %v\n", iface == nil)
	iface = nilInterface()
	fmt.Printf("nil-init-interface == nil: %v\n", iface == nil)
	iface = (IAlive)(nil)
	fmt.Printf("explicit-nil-init-interface == nil: %v\n", iface == nil)
	iface = (*Person)(nil)
	fmt.Printf("typed-nil-init-interface == nil: %v\n", iface == nil)
	fmt.Printf("typed-nil-init-interface == interface-nil: %v\n", iface == IAlive(nil))
	fmt.Printf("typed-nil-init-interface == typed-nil: %v\n", iface == (*Person)(nil))
}

func nilSlice() []int {
	var nilSlice = []int(nil)
	return nilSlice
}

func nilMap() map[string]int {
	var nilMap = map[string]int(nil)
	return nilMap
}

func nilBasicPointer() *int {
	var nilPointer *int = nil
	return nilPointer
}

func nilStructPointer() *Person {
	var nilPointer *Person = nil
	return nilPointer
}

func nilInterfacePointer() *IAlive {
	var nilPointer *IAlive = nil
	return nilPointer
}

func nilInterface() IAlive {
	var nilPointer IAlive = nil
	return nilPointer
}
