package main

import (
	"errors"
	"fmt"
)

func main() {
	// defer
	defer finish()
	defer globalErrorHandler()
	fmt.Println("Program has been started")
	fmt.Println("Program is working")

	// handle error
	q, err := divide2(1, 0)
	if err != nil {
		fmt.Println("divide2 error: ", err)
	} else {
		fmt.Println(q)
	}

	// panic
	fmt.Println(divide(15, 5))
	fmt.Println(divide(4, 0)) // <-- panic
}

func finish() {
	fmt.Println("Program has been finished")
}

func divide(x, y float64) float64 {
	if y == 0 {
		panic("Division by zero!")
	}
	return x / y
}

func divide2(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Division by zero!")
	}
	return x / y, nil
}

func globalErrorHandler() {
	if e := recover(); e != nil {
		fmt.Println("Error: ", e)
	}
}
