package main

import (
	"artem/functions/temperature"
	"errors"
	"fmt"
	"math/big"
)

type Point struct {
	x int
	y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Predicate func(int) bool
type Operation func(int, int) int

func main() {
	// методы для кастомных типов
	var c temperature.Celcius = +101
	fmt.Printf("c.ToFarenheite(): %0.1f\n", c.ToFarenheite())

	//var t int = +100
	//fmt.Println(t.ToFarenheite()) // error: на базовом типе метод кастомного типа вызывать нельзя

	var f temperature.Farenheite = -10
	fmt.Printf("f.ToCelcius(): %0.1f\n", f.ToCelcius())

	fmt.Println("---------------------------------")

	// рекурсивные функции
	var n uint = 10
	fmt.Printf("fib(%d)        = %d\n", n, fib(n))
	fibN := fibDynamic(n)
	fmt.Printf("fibDynamic(%d) = %v\n", n, fibN.String())
	// fibDynamic(1_000_000) // stackoverflow

	fmt.Println("---------------------------------")

	// вызов функций с параметрами
	sum(2, 1)
	subtract(2, 1)

	// передача параметров по значению
	x := 1
	fmt.Printf("Before inc: %d\n", x)
	inc(x)
	fmt.Printf("After inc: %d\n", x)

	p := Point{x: 1, y: 2}
	fmt.Printf("Before move: %s\n", p.String())
	move(p, 10, 10)
	fmt.Printf("After move: %s\n", p.String())

	fmt.Println("---------------------------------")

	// передача параметров по ссылке
	fmt.Printf("Before inc(*): %d\n", x)
	incRef(&x)
	fmt.Printf("After inc(*): %d\n", x)

	fmt.Printf("Before move(*): %s\n", p.String())
	moveRef(&p, 10, 10)
	fmt.Printf("After move(*): %s\n", p.String())

	fmt.Println("---------------------------------")

	// возврат результата функции по ссылке
	fmt.Printf("Before scale()*: %s\n", p.String())
	sp := scale(&p, 2)
	fmt.Printf("After scale()*: %s\n", sp.String())
	p.x = 50
	p.y = 50
	fmt.Printf("After scaled point change: %s\n", sp.String())

	fmt.Println("---------------------------------")

	// переменное число параметров функции
	add(1, 2, 3, 4)
	add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	numbers := []int{1, 2, 3}
	add(numbers...)
	add([]int{1, 2, 3, 4}...)

	fmt.Println("---------------------------------")

	// несколько возвращаемых значений
	n1, n2 := 10, 4
	q, r := divide(n1, n2)
	fmt.Printf("divide(%d, %d): q=%d, r=%d\n", n1, n2, q, r)

	fmt.Println("---------------------------------")

	// анонимная функция
	var mult Operation = func(x int, y int) int { return x * y }
	fmt.Printf("multiply(%d, %d): %d\n", n1, n2, mult(n1, n2))

	// функция как возвращаемое значение другой функции
	sub, _ := getOperation("-")
	fmt.Printf("subtract(%d, %d): %d\n", n1, n2, sub(n1, n2))

	// передача функции параметром в другую функцию
	sumPositive := sumBy(isPositive, -1, 2, -3, 4, -5, 6)
	fmt.Printf("sum(positive): %d\n", sumPositive)

	// замыкание окружения
	squareNext := square(2)
	fmt.Printf("%d, %d, %d, %d, %d", squareNext(), squareNext(), squareNext(), squareNext(), squareNext())
}

func fib(n uint) uint {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

var fibcache = map[uint]big.Int{
	1: *big.NewInt(int64(1)),
	2: *big.NewInt(int64(1)),
}

func fibDynamic(n uint) big.Int {
	if f, ok := fibcache[n]; ok {
		return f
	} else {
		fn1, fn2 := fibDynamic(n-1), fibDynamic(n-2)
		f = *new(big.Int).Add(&fn1, &fn2)
		fibcache[n] = f
		return f
	}
}

func sum(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x int, y int) int { return x * y }

func inc(x int) int {
	x++
	return x
}

func incRef(x *int) int {
	(*x)++
	return *x
}

func move(p Point, dx int, dy int) Point {
	p.x += dx
	p.y += dy
	return p
}

func moveRef(p *Point, dx int, dy int) Point {
	p.x += dx
	p.y += dy
	return *p
}

func scale(p *Point, k int) *Point {
	p.x *= k
	p.y *= k
	return p
}

func add(numbers ...int) int {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// именованные возвращаемые значения
func divide(x, y int) (q int, r int) {
	q = x / y
	r = x % y
	return
}

func isPositive(n int) bool {
	return n > 0
}

func sumBy(criteria Predicate, numbers ...int) int {
	s := 0
	for _, num := range numbers {
		if criteria(num) {
			s += num
		}
	}
	return s
}

func getOperation(op string) (operation Operation, err error) {
	switch op {
	case "+":
		return sum, nil
	case "-":
		return subtract, nil
	case "*":
		return multiply, nil
	case "/":
		return func(x int, y int) int { return x / y }, nil
	default:
		return nil, errors.New("unknown operation")
	}
}

func square(origin int) func() int {
	var s int = origin
	return func() int {
		result := s * s
		s++
		return result
	}
}
