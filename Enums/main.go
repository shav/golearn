package main

import (
	"artem/enums/weekdays"
	"fmt"
)

func main() {
	day1, _ := weekdays.FromString("Monday")
	fmt.Println(day1)

	day2 := weekdays.Sunday
	fmt.Println(day2)

	fmt.Println("---------------------------------")
}
