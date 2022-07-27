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

	fmt.Println("--------------------------------------")
}
