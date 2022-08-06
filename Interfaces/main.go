package main

import (
	"fmt"
)

type IMovable interface {
	Move()
}

type IAlive interface {
	Eat()
	SetName(name string)
}

type INaming interface {
	SetName2(name string)
}

type ISleeper interface {
	Sleep()
}

type IAnimal interface {
	IAlive
	ISleeper
	IMovable
}

type Animal struct {
	Name string
}

func (animal Animal) Eat() {
	fmt.Printf("%s eats\n", animal.Name)
}

func (animal Animal) Sleep() {
	fmt.Printf("%s sleeps\n", animal.Name)
}

func (animal Animal) Move() {
	fmt.Printf("%s moves\n", animal.Name)
}

func (animal Animal) SetName(name string) {
	animal.Name = name
}

func (animal *Animal) SetName2(name string) {
	animal.Name = name
}

type Dog struct {
	Animal
}

func (dog Dog) Bark() {
	fmt.Printf("%s barks\n", dog.Name)
}

type Cat struct {
	Animal
}

func (cat Cat) Meow() {
	fmt.Printf("%s meows\n", cat.Name)
}

func (cat Cat) Sleep() {
	fmt.Println("Cat zzzzzz...")
	cat.Animal.Sleep()
}

type Plane struct {
	Model string
}

func (plane Plane) Move() {
	fmt.Printf("%s fly\n", plane.Model)
}

// -------------------------------------------------------------

type IConflict1 interface {
	Read()
}

type IConflict2 interface {
	Read(name string)
	Write(name string)
}

type ConflictImpl struct {
	name string
}

func (c *ConflictImpl) Read() {
}

// Перегрузка методов не поддерживается!
//func (c *ConflictImpl) Read(name string) {
//}

func (c *ConflictImpl) Write(name string) {
}

// -------------------------------------------------------------

type ConflictBase1 struct {
}

func (c *ConflictBase1) Print() {
	fmt.Println("ConflictBase1")
}

type ConflictBase2 struct {
}

func (c *ConflictBase2) Print() {
	fmt.Println("ConflictBase2")
}

type Conflict struct {
	ConflictBase1
	ConflictBase2
}

func (c *Conflict) Print() {
	c.ConflictBase1.Print()
	c.ConflictBase2.Print()
	fmt.Println("Conflict")
}

func main() {
	var cat = Cat{Animal{Name: "Murka"}}
	var dog = Dog{Animal{Name: "Rex"}}
	var airbus = Plane{Model: "Airbus"}

	// Вызов методов интерфейса
	var animal IAnimal = cat
	animal.Sleep()

	animal = dog
	animal.Sleep()
	// animal.Bark() // compile error: На этапе компиляции неизвестно, что по факту в переменной animal хранится Dog

	// animal = airbus // error: Plane не реализует интерфейс IAnimal

	// Так не работает - объект передаётся в метод по значению
	cat.SetName("MURKA")
	fmt.Println(cat)

	// Так тоже не работает, не смотря на то что метод вызывается через указатель -
	// всё-равно он разыменовывается и передаётся в метод по значению.
	pCat := &cat
	pCat.SetName("MURKA")
	fmt.Println(cat)
	// Для этого нужно реализовывать интерефейс с модифицирующими методами через указатели!
	var namedAnimal INaming = &cat
	namedAnimal.SetName2("MURKA2")
	fmt.Println(cat)

	fmt.Println("--------------------------------------")

	// Присваивание интерфейсных переменных
	var alive IAlive = animal
	alive.Eat()

	var sleeper ISleeper = dog
	sleeper.Sleep()

	// error: интерфейсы ISleeper и IAlive несовместимы между собой, хоть и по факту в переменной alive содержится объект, реализующий ISleeper
	// sleeper = alive
	// Здесь аналогично:
	// animal = alive
	// Нужно в runtime проверять через утверждение типа, реализует ли конкретный объект заданный интерфейс
	if animal, ok := alive.(IAnimal); ok {
		animal.Move()
	}

	fmt.Println("--------------------------------------")

	// Проверка на равенство
	var animal1 IAnimal = dog
	var animal2 IAnimal = dog
	var animal3 IAnimal = cat
	fmt.Printf("interface == interface (equals): %v\n", animal1 == animal2)
	fmt.Printf("interface == interface (not equals): %v\n", animal1 == animal3)
	fmt.Printf("interface == object (equals): %v\n", animal1 == dog)
	fmt.Printf("interface == object (not equals): %v\n", animal1 == cat)

	fmt.Println("--------------------------------------")

	// Полиморфизм
	movers := []IMovable{cat, dog, airbus}
	for _, mover := range movers {
		mover.Move()
	}

	fmt.Println("--------------------------------------")

	// Ковариантность/контравариантность
	animals := []IAnimal{cat, dog}
	// movers = animals // error: из коробки ковариантность не работает
	fmt.Println(animals)

	fmt.Println("--------------------------------------")

	// Конфликты при имплементации интерфейсов
	var con1 IConflict1 = &ConflictImpl{name: "Qwerty"}
	fmt.Println(con1)

	// Не получается реализовать метод из IConflict2,
	// т.к. метод с таким же именем существует в IConflict1, а перегрузка методов не поддерживается
	//var con2 IConflict2 = &ConflictImpl{name: "Qwerty"}

	fmt.Println("--------------------------------------")

	// Конфликты при "наследовании"-композиции
	var con Conflict = Conflict{}
	con.Print()
}
