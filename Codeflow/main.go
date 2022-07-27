package main

import (
	"fmt"
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

	var pow = 1
	for pow <= 1024 {
		fmt.Printf("%d, ", pow)
		pow *= 2
	}
	fmt.Println()

	rand.Seed(time.Now().Unix())
	var degrees = 0
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

	fmt.Println("--------------------------------------")

	// Switch
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Сегодня понедельник.")
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
}
