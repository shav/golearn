package main

import (
	"fmt"
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
}
