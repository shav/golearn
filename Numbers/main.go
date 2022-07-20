package main

import (
	"fmt"
	"math"
)

func main() {
	// Целые числа
	var int0 int
	var int320 int32
	var int640 int64
	var uint320 uint32
	var uint640 uint64

	int320 = 1 << 30
	int0 = int(int320) // повышающее преобразование
	fmt.Printf("int32 -> int: %d\n", int0)

	int640 = 1 << 31
	int320 = int32(int640) // понижающее преобразование
	fmt.Printf("int64 -> int32: %d\n", int320)

	int320 = 1 << 30
	int640 = int64(int320) // повышающее преобразование
	fmt.Printf("int32 -> int64: %d\n", int640)

	// int320 = -(1 << 30)
	// uint640 = int(int320) // несовместимое преобразование
	// fmt.Println(uint640)

	uint640 = 1 << 63
	int320 = int32(uint640) // понижающее преобразование
	fmt.Printf("uint -> int32: %d\n", int320)

	uint640 = 1 << 32
	uint320 = uint32(uint640) // понижающее преобразование
	fmt.Printf("uint64 -> uint32: %d\n", uint320)

	int0 = 1
	int320 = 1
	fmt.Println(int0 == int(int320))
	fmt.Println(int0 < int(int320))

	sum := int0 + int(int320)
	fmt.Printf("sum = %d \n", sum)

	int640 = 1 << 62
	var int641 int64 = 1 << 62
	sumOverflow := int640 + int641
	fmt.Printf("sumOverflow = %d \n", sumOverflow)
	multiOverflow := int640 * int641
	fmt.Printf("multiOverflow = %d \n", multiOverflow)

	d1 := 3
	d2 := 2
	fmt.Printf("+int / +int: %[1]d = %[3]d * %[2]d %+[4]d\n", d1, d2, d1/d2, d1%d2)
	fmt.Printf("-int / -int: %[1]d = %[3]d * %[2]d %+[4]d\n", -d1, -d2, -d1/-d2, -d1%-d2)
	fmt.Printf("-int / +int: %[1]d = %[3]d * %[2]d %+[4]d\n", -d1, d2, -d1/d2, -d1%d2)
	fmt.Printf("+int / -int: %[1]d = %[3]d * %[2]d %+[4]d\n", d1, -d2, d1/-d2, d1%-d2)

	// // Error:
	// var zero int = 0
	// fmt.Printf("div by zero: %d\n", 1/zero)

	fmt.Println("--------------------------------------")

	// Вещественные числа
	var f0 float32 = 3.14
	var f1 float32 = 3.1401
	var f2 float32 = 3.5001
	fmt.Println(f0 == f1)
	fmt.Println(f0 < f1)
	fmt.Println(f0 + f1)
	fmt.Println(f0 * f1)
	fmt.Println(10.0 / 4)
	fmt.Println(10.0 / 3.14)
	fmt.Printf("float -> int: %d\n", int(f2))
	fmt.Printf("round(float): %d\n", int(math.Round(float64(f2))))

	fmt.Println("--------------------------------------")

	// Комплексные числа
	var c0 complex64 = 1 + 2i
	var c1 complex64 = 1 + 2i
	fmt.Println(c0 == c1)
	//fmt.Println(c0 < c1) // unsupported
	fmt.Println(c0 + c1)
	fmt.Println(c0 * c1)
	fmt.Println(c0 / c1)
}
