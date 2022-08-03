package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Условия
	var year = 2100
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("%d year is leap\n", year)
	} else {
		fmt.Printf("%d year is not leap\n", year)
	}

	dayOfWeek := time.Now().Weekday()
	if dayOfWeek == time.Sunday {
		fmt.Println("Sunday")
	} else if dayOfWeek == time.Saturday {
		fmt.Println("Saturday")
	} else {
		fmt.Println("Workday")
	}

	fmt.Println("--------------------------------------")

	// Циклы
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d, ", i)
	}
	fmt.Println()

	// while
	var pow = 1
	for pow <= 1024 {
		fmt.Printf("%d, ", pow)
		pow *= 2
	}
	fmt.Println()

	rand.Seed(time.Now().Unix())
	var degrees = 0
	// бесконечный цикл
	for {
		fmt.Printf("%d, ", degrees)
		degrees++
		if degrees >= 360 {
			degrees = 0
		}
		if rand.Intn(10) == 0 {
			break
		}
	}
	fmt.Println()

	// foreach
	var users = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users {
		fmt.Println(index, value)
	}
	for _, value := range users {
		fmt.Println(value)
	}

	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	for _, num := range numbers {
		if num < 0 {
			continue // переходим к следующей итерации
		}
		if num > 6 {
			break // выходим из цикла
		}
		fmt.Printf("%d, ", num)
	}
	fmt.Println()

	fmt.Println("--------------------------------------")

	// Вложенный цикл
outerLoopLabel:
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j == 5 {
				continue
			}
			if j == i {
				break
			}
			if i*j > 63 {
				break outerLoopLabel
			}
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("--------------------------------------")

	// Switch
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Сегодня понедельник.")
		// break добавляется автоматически, в отличие от других ЯП
	case time.Tuesday:
		fmt.Println("Сегодня вторник.")
	case time.Wednesday:
		fmt.Println("Сегодня среда.")
	case time.Thursday:
		fmt.Println("Сегодня четверг.")
	case time.Friday:
		fmt.Println("Сегодня пятница.")
	case time.Saturday:
		fmt.Println("Сегодня суббота.")
	case time.Sunday:
		fmt.Println("Сегодня воскресенье.")
	}

	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Сегодня будний день")
	case time.Saturday, time.Sunday:
		fmt.Println("Сегодня выходной день")
	}

	size := "XXL"
	switch size {
	case "XXS":
		fmt.Println("очень очень маленький")
	case "XS":
		fmt.Println("очень маленький")
	case "S":
		fmt.Println("маленький")
	case "M":
		fmt.Println("средний")
	case "L":
		fmt.Println("большой")
	case "XL":
		fmt.Println("очень большой")
	case "XXL":
		fmt.Println("очень очень большой")
	default:
		fmt.Println("неизвестно")
	}

	str := "a b c\td\nefg hi"
	for _, char := range str {
		switch char {
		case ' ', '\t', '\n':
			break
		default:
			fmt.Printf("%c", char)
		}
	}
	fmt.Println()

	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("AM")
	default:
		fmt.Println("PM")
	}

	// Значения case-веток не обязаны быть константами!
	binary := 16
	switch binary {
	case pow2(1):
		fmt.Println("pow2(1)")
	case pow2(2):
		fmt.Println("pow2(2)")
	case pow2(3):
		fmt.Println("pow2(3)")
	case pow2(4):
		fmt.Println("pow2(4)")
	default:
		fmt.Println("pow2(undefined)")
	}

	nextStop := "B"
	fmt.Print("Stops ahead of us: ")
	switch nextStop {
	case "A":
		fmt.Print("A->")
		fallthrough
	case "B":
		fmt.Print("B->")
		fallthrough
	case "C":
		fmt.Print("C->")
		fallthrough
	case "D":
		fmt.Print("D->")
		fallthrough
	case "E":
		fmt.Print("E")
	}
	fmt.Println()

	var data interface{}
	data = 3.14
	switch data.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("unknown")
	}

	str = "abc_z_def"
ForLabel:
	for _, char := range str {
		switch char {
		case 'a':
			fmt.Print("A")
		case 'b':
			fmt.Print("B")
		case 'c':
			fmt.Print("C")
		case 'd':
			fmt.Print("D")
		case 'z':
			fmt.Print("Z")
			break ForLabel // выход из цикла (просто break выполнит выход из case-ветки)
		}
	}
}

func pow2(num int) int {
	return int(math.Pow(2, float64(num)))
}
