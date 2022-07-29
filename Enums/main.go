package main

import (
	"artem/enums/seasons"
	"artem/enums/weekdays"
	"fmt"
)

func main() {
	day1, _ := weekdays.FromString("Monday")
	fmt.Println(day1)

	day2 := weekdays.Sunday
	fmt.Println(day2)

	fmt.Println("---------------------------------")

	season1, _ := seasons.FromString("winter")
	fmt.Println(season1)

	season2 := seasons.Summer
	fmt.Println(season2)
}
