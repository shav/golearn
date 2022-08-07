package main

import (
	"artem/collections/mc"
	"fmt"
)

type FlagsType struct {
	F0 mc.Flag
	F1 mc.Flag
	F2 mc.Flag
}

var Flags = FlagsType{
	F0: 1 << 0,
	F1: 1 << 1,
	F2: 1 << 2,
}

func main() {
	// Набор битов
	bits := mc.BitSet(0)
	bits.Set(Flags.F2 | Flags.F1)
	fmt.Printf("%b\n", bits)
	bits.Toggle(Flags.F1)
	fmt.Printf("%b\n", bits)
	bits.Clear(Flags.F2)
	fmt.Printf("%b\n", bits)

	fmt.Println("--------------------------------------")

	set := mc.NewSet[int]()
	set.Add(1, 2, 3)
	set.Add(1, 2)
	fmt.Println(set)

	set.Remove(1)
	fmt.Println(set)

	searchValue := 2
	fmt.Printf("set contains %d: %v\n", searchValue, set.Contains(searchValue))
	searchValue = 1
	fmt.Printf("set contains %d: %v\n", searchValue, set.Contains(searchValue))

	fmt.Println("--------------------------------------")

	// Стек
	stack := mc.NewEmptyStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)

	stackTop, err := stack.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("stack top: %d\n", stackTop)
	}

	stackTop, err = stack.Pop()
	stackTop, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("stack top: %d\n", stackTop)
	}

	searchValue = 2
	fmt.Printf("stack contains %d: %v\n", searchValue, stack.Contains(searchValue))
	searchValue = 1
	fmt.Printf("stack contains %d: %v\n", searchValue, stack.Contains(searchValue))

	stack.Clear()
	fmt.Println(stack)

	stackTop, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("stack top: %d\n", stackTop)
	}

	fmt.Println("--------------------------------------")

	// Список (на основе массивов)
	list := mc.NewEmptyArrayList[int]()
	list.Add(10, 20, 30, 20, 40)
	fmt.Println(list)
	fmt.Printf("list length: %d\n", list.Length())

	list.Remove(20)
	fmt.Println(list)

	listItem, err := list.GetByIndex(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("list[1]: %d\n", listItem)
	}
	listItem, err = list.GetByIndex(1000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("list item: %d\n", listItem)
	}

	searchValue = 10
	fmt.Printf("list contains %d: %v\n", searchValue, list.Contains(searchValue))
	searchValue = 20
	fmt.Printf("list contains %d: %v\n", searchValue, list.Contains(searchValue))

	list.Clear()
	fmt.Println(list)
}
