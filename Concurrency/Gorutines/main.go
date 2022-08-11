package main

import (
	"fmt"
)

func main() {
	go printMessage("Hello")
	go printMessage("World")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Printf("------------------------------")
}

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
	}
}
