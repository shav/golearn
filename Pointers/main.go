package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
	age  int
}

type ITalker interface {
	talk() string
}

func shout(t ITalker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func main() {
	var num1 int = 1
	// Инициализация указателя (взятие адреса)
	var pNum *int = &num1
	fmt.Printf("%v (%[1]T): %d\n", pNum, *pNum)
	// Изменение значения объекта, на который указывает указатель
	*pNum = 11
	fmt.Println(num1)

	var num2 int = 2
	pNum = &num2
	fmt.Printf("%v (%[1]T): %d\n", pNum, *pNum)
	*pNum = 22
	fmt.Println(num2)

	fmt.Println("--------------------------------------")

	// Разыменование указателя
	// При разыменовании создаётся копия объекта, дальнейшие изменения через указатель на неё уже не влияют
	var num22 = *pNum
	*pNum = 220
	fmt.Println(num2)
	fmt.Println(num22)

	fmt.Println("--------------------------------------")

	// работа с переменными-указателями
	var pNum2 *int = pNum
	*pNum2 = 222
	fmt.Println(num2)
	pNum2 = &num1
	*pNum2 = 111
	fmt.Println(num1)
	fmt.Println(num2)

	fmt.Println("--------------------------------------")

	// Указатель на указатель
	var ppNum **int = &pNum
	fmt.Printf("%v (%[1]T): 0x%x(%[2]T): %d\n", ppNum, *ppNum, **ppNum)
	// *ppNum += 16 // error: адресная арифметика не поддерживается

	fmt.Println("--------------------------------------")

	// Несоответствие типов указателей
	var str string = "abc"
	var pStr *string = &str
	fmt.Printf("%v (%[1]T): %s\n", pStr, *pStr)
	// pStr = &num1 // error
	// pStr = pNum // error

	fmt.Println("--------------------------------------")

	// Указатели на массивы
	var board = [2][2]int{{1, 0}, {0, 1}}
	fmt.Println(board)
	var pBoard *[2][2]int = &board
	pBoard[0][0] = 2
	fmt.Println(board)
	(*pBoard)[0][0] = 5
	fmt.Println(board)

	fmt.Println("--------------------------------------")

	// Указатели на структуры
	var person = Person{name: "Artem", age: 31}
	var pPerson = &person
	fmt.Println(person)

	pPerson.age = 20
	fmt.Println(person)

	var pName = &person.name
	*pName = "ARTEM"
	fmt.Println(person)

	pPerson3 := &Person{name: "Vovan", age: 29}
	fmt.Println(*pPerson3)

	pPerson4 := new(Person)
	pPerson4.name = "Dimon"
	pPerson4.age = 10
	fmt.Println(*pPerson4)

	var pPerson2 = &person
	fmt.Printf("*struct == *struct (pointers to the same object): %v\n", pPerson2 == pPerson)
	fmt.Printf("*struct == *struct (pointers to different objects): %v\n", pPerson3 == pPerson)

	fmt.Println("--------------------------------------")

	// Указатели на интерфейсы
	var nack = martian{}
	shout(nack)
	shout(&nack)

	pew := laser(2)
	//shout(pew) // error: интерфейс ITalker в типе laser реализован через указатель
	shout(&pew)

	fmt.Println("--------------------------------------")

	// Передача аргумента в функцию по ссылке
	fmt.Printf("int: %d\n", num1)
	increment(num1)
	fmt.Printf("increment(int): %d\n", num1)

	incrementRef(&num1)
	fmt.Printf("increment(*int): %d\n", num1)

	pNum1 := getRef(&num1)
	*pNum1 += 10
	fmt.Printf("return *int: %d\n", num1)
}

func increment(num int) int {
	num++
	return num
}

func incrementRef(num *int) int {
	*num++
	return *num
}

func getRef(num *int) *int {
	return num
}
