package main

import (
	"artem/enums/roles"
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

	fmt.Println("---------------------------------")

	role1, _ := roles.FromString("Admin")
	fmt.Println(role1)

	role2 := roles.Guest
	fmt.Println(role2)
}
