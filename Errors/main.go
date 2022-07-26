package main

import (
	"errors"
	"fmt"
	goerrors "github.com/go-errors/errors"
	"io"
	"runtime"
)

type IndexOutOfRangeError struct {
	message string
	index   int
}

func (error IndexOutOfRangeError) Error() string {
	return error.message
}

type ArgumentError struct {
	message string
	value   any
}

func (error ArgumentError) Error() string {
	return error.message
}

func main() {
	// Отложенный вызов функциий в конце программы
	defer finish()
	defer globalErrorHandler()
	fmt.Println("Program has been started")
	fmt.Println("Program is working")

	fmt.Println("---------------------------------")

	// Обработка ошибок
	q, err := divide2(1, 0)
	if err != nil {
		fmt.Println("divide2 error: ", err)
	} else {
		fmt.Println(q)
	}

	fmt.Println("---------------------------------")

	// Кастомные типы ошибок
	err = raiseError(0)
	if err != nil {
		// Проверка типа ошибки (вариант 1):
		if outOfRangeError, ok := err.(IndexOutOfRangeError); ok {
			fmt.Printf("%s: %d\n", outOfRangeError.message, outOfRangeError.index)
		}
		if argumentError, ok := err.(ArgumentError); ok {
			fmt.Printf("%s: %v\n", argumentError.message, argumentError.value)
		}

		// Проверка типа ошибки (вариант 2):
		switch error := err.(type) {
		case IndexOutOfRangeError:
			fmt.Printf("%s: %d\n", error.message, error.index)
		case ArgumentError:
			fmt.Printf("%s: %v\n", error.message, error.value)
		}

		// Проверка ошибки (ошибка должна быть известной константой!)
		errors.Is(err, io.EOF)
	}

	fmt.Println("---------------------------------")

	tryCatched := TryCatch()
	fmt.Println("try/catch: ", tryCatched)
	fmt.Println()

	tryCatched2 := TryCatchWithResult()
	fmt.Println("try/catch with result: ", tryCatched2)

	fmt.Println("---------------------------------")

	// Паника
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

		// Получить стектрейс
		var stackTrace = make([]byte, 1000 /* TODO: Пока непонятно, как заранее вычислить размер стектрейса */)
		runtime.Stack(stackTrace, true)
		fmt.Println(string(stackTrace))
		// Вариант 2 (с использованием стороннего пакета)
		fmt.Println(goerrors.Wrap(e, 2).ErrorStack())

		//debug.PrintStack()
	}
}

func raiseError(code int) error {
	switch code {
	case 0:
		return IndexOutOfRangeError{message: "Index is out of range", index: 13}
	case 1:
		return ArgumentError{message: "Wrong argument value", value: "777"}
	}
	return nil
}

func TryCatch() int {
	defer func() {
		// Имитация try/catch:
		// После восстановления выполнение кода текущей функции ниже ошибки прерывается,
		// но выполнение вызывающего внешнего кода продолжается!
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Continue...")
		}
	}()
	result := 1
	panic("TryCatch: fail")
	result = 2
	return result
}

func TryCatchWithResult() (result int) {
	defer func() {
		// Имитация try/catch:
		// После восстановления выполнение кода текущей функции ниже ошибки прерывается,
		// но выполнение вызывающего внешнего кода продолжается!
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Continue...")
			// Данное значение вернётся наружу в вызывающий внешний код
			result = 10
		}
	}()
	result = 1
	panic("TryCatchWithResult: fail")
	result = 2
	return result
}
