package main

import (
	"fmt"
)

func main() {
	// Константы
	const cNum0 int = 100
	fmt.Println(cNum0)

	//cNum = 101 // error: значение константы изменять нельзя

	const cNum00 int = cNum0
	fmt.Println(cNum00)

	//var numVar int = 1
	//const cNum3 int = numVar // error: константу нельзя инициализировать значением переменной

	const (
		cNum1 = 40
		cNum2 = 50
	)
	fmt.Println(cNum1)
	fmt.Println(cNum2)

	const (
		_ = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Printf("Kb = %d bytes\n", kb)
	fmt.Printf("Mb = %d bytes\n", mb)
	fmt.Printf("Gb = %d bytes\n", gb)
	fmt.Printf("Tb = %d bytes\n", tb)

	const (
		flag1 = 1 << iota
		flag2
		flag3
		flag4
		flag5
	)
	fmt.Printf("flags = %d, %d, %d, %d, %d\n", flag1, flag2, flag3, flag4, flag5)
	fmt.Printf("flag1 | flag2 | flag3 = %d\n", flag1|flag2|flag3)

	const cBigNum1 int64 = 1e+18
	fmt.Printf("Big const: %d\n", cBigNum1)

	// Константы являются нетипизированными и могут быть безразмерными
	const cBigNum2 = 100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
	const cBigNum3 = 100000000000000000000000000000000000000000000
	var div float64 = cBigNum2 / cBigNum3
	fmt.Println(div)

	// const n int64 = 9876543210 * 9876543210 // error: overflow
	// fmt.Println(n)
}
