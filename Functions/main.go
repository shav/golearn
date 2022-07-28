package main

import (
	"artem/functions/temperature"
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

	// передача параметров по ссылке
	fmt.Printf("Before inc(ref): %d\n", x)
	incRef(&x)
	fmt.Printf("After inc(ref): %d\n", x)

	fmt.Printf("Before move(ref): %s\n", p.String())
	moveRef(&p, 10, 10)
	fmt.Printf("After move(ref): %s\n", p.String())

	fmt.Println("---------------------------------")
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
