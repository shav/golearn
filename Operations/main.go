package main

import (
	"fmt"
)

func main() {
	// арифметические операции
	num1 := 4
	num2 := 2
	fmt.Printf("int + int: %d\n", num1+num2)
	fmt.Printf("int - int: %d\n", num1-num2)
	fmt.Printf("int * int: %d\n", num1*num2)
	fmt.Printf("int / int: %d\n", num1/num2)
	fmt.Printf("int %% int: %d\n", num1%num2)

	num1++
	fmt.Printf("int++: %d\n", num1)
	// ++num1 // error: префиксного инкремента нет

	num1--
	fmt.Printf("int--: %d\n", num1)
	// --num1 // error: аналогично, префиксного декремента тоже нет

	// Инкремент/декремент не может участвовать в выражениях, только сам по себе
	// a:=2*num1++
	// array :=[]int{1,2,3}
	// fmt.Println(array[num1++])

	fmt.Println("--------------------------------------")

	// сравнение чисел
	num1 = 4
	num2 = 2
	fmt.Printf("int == int: %v\n", num1 == num2)
	fmt.Printf("int < int: %v\n", num1 < num2)
	fmt.Printf("int <= int: %v\n", num1 <= num2)
	fmt.Printf("int > int: %v\n", num1 > num2)
	fmt.Printf("int >= int: %v\n", num1 >= num2)
	fmt.Printf("int != int: %v\n", num1 != num2)

	fmt.Println("--------------------------------------")

	// битовые операции
	num := 16
	num1 = 5
	num2 = 3
	fmt.Printf("%b << 2: %b (=%[2]d)\n", num, num<<2)
	fmt.Printf("%b >> 2: %b (=%[2]d)\n", num, num>>2)
	fmt.Printf("%b & %b: %b (=%[3]d)\n", num1, num2, num1&num2)
	fmt.Printf("%b | %b: %b (=%[3]d)\n", num1, num2, num1|num2)
	fmt.Printf("%b ^ %b: %b (=%[3]d)\n", num1, num2, num1^num2)
	fmt.Printf("%b &^ %b: %b (=%[3]d)\n", num1, num2, num1&^num2)

	fmt.Println("--------------------------------------")

	// конкатенация строк
	str1 := "str1"
	str2 := "str2"
	fmt.Printf("str + str: %s\n", str1+str2)

	fmt.Println("--------------------------------------")

	// сравнение строк
	abc := "abc"
	ABC := "ABC"
	abcd := "abcd"
	fmt.Printf("%s == %s: %v\n", abc, abc, abc == abc)
	fmt.Printf("%s == %s: %v\n", abc, abcd, abc == abcd)
	fmt.Printf("%s == %s: %v\n", abc, ABC, abc == ABC)
	fmt.Printf("%s < %s: %v\n", abc, abcd, abc < abcd)
	fmt.Printf("%s > %s: %v\n", abc, abcd, abc > abcd)
	fmt.Printf("%s != %s: %v\n", abc, abc, abc != abc)
	fmt.Printf("%s != %s: %v\n", abc, abcd, abc != abcd)
	fmt.Printf("%s != %s: %v\n", abc, ABC, abc != ABC)

	fmt.Println("--------------------------------------")

	// логические операции
	bool1 := true
	bool2 := false
	fmt.Printf("bool && bool: %v\n", bool1 && bool2)
	fmt.Printf("bool || bool: %v\n", bool1 || bool2)
	fmt.Printf("!bool: %v\n", !bool1)
}
