package main

import (
	"fmt"
)

func main() {
	defer finish()
	fmt.Println("Program has been started")
	fmt.Println("Program is working")
}

func finish() {
	fmt.Println("Program has been finished")
}
