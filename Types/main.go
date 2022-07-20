package main

import (
	"fmt"
	"types/temperature"
)

func main() {
	// Целочисленные типы
	var num0 byte = 1
	//num0 = 1000 // error: overflow
	fmt.Println(num0)

	var num1 int = -(2 << 61)
	fmt.Println(num1)

	var num2 int32 = -2 << 30
	fmt.Println(num2)

	var num3 int64 = -2 << 62
	fmt.Println(num3)

	var num4 uint = 2 << 61
	fmt.Println(num4)

	var num5 uint32 = 2 << 30
	fmt.Println(num5)

	var num6 uint64 = 2 << 62
	fmt.Println(num6)

	var numOct int = 0567
	//numOct = 0569 // error
	fmt.Println(numOct)

	var numBin int = 0b1010
	fmt.Println(numBin)

	var numHex int = 0x12ABC
	fmt.Println(numHex)

	fmt.Println("------------------------------------------------")

	// Вещественные числа
	var num7 float32 = 23.00121323254235765546834690
	fmt.Println(num7)

	var num71 float32 = 2e+10
	fmt.Println(num71)

	var num8 float64 = 376123457862892137091285709168046912359023239859038590237568921346042769023769237689012370698.7657655345123456365231334
	fmt.Println(num8)

	var num81 float64 = -2e-100
	fmt.Println(num81)

	var num82 float64 = 2e+307
	fmt.Println(num82)

	fmt.Println("------------------------------------------------")

	// Комплексные числа
	var complex0 complex64 = 0
	fmt.Println(complex0)

	var complex1 complex64 = 1 - 2i
	fmt.Println(complex1)

	fmt.Println("------------------------------------------------")

	// Булевские значения
	var f0 bool = false
	fmt.Println(f0)

	var f1 bool = false && true
	fmt.Println(f1)

	var t0 bool = true
	fmt.Println(t0)

	var t1 bool = false || true
	fmt.Println(t1)

	fmt.Println("------------------------------------------------")

	// Строки
	var str0 string = ""
	fmt.Println(str0)

	var str1 string = "Hello"
	fmt.Println(str1)

	var str2 string = "\t\"Hello\\Artem\""
	fmt.Println(str2)

	fmt.Println("------------------------------------------------")

	// Символы
	var char0 = 'a'
	fmt.Printf("%c\n", char0)

	var char1 = '"'
	fmt.Printf("%c\n", char1)

	var char2 = rune(223)
	fmt.Printf("%c\n", char2)

	var char3 = rune(63916)
	fmt.Printf("%c\n", char3)

	fmt.Println("------------------------------------------------")

	// Кастомные типы на основе базовых
	var c temperature.Celcius = +101
	fmt.Printf("c: %0.1f\n", c)
	fmt.Printf("rune(c): %d\n", rune(c))
	fmt.Printf("string(c): %s\n", string(int(c)))
	fmt.Printf("c.ToFarenheite(): %0.1f\n", c.ToFarenheite())
	var f temperature.Farenheite = -10
	fmt.Printf("f: %0.1f\n", f)
	fmt.Printf("int(f): %d\n", int(f))
	fmt.Printf("f.ToCelcius(): %0.1f\n", f.ToCelcius())
}
