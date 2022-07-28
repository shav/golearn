package main

import (
	"artem/functions/temperature"
	"fmt"
	"math/big"
)

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
	var n uint = 100_000
	//fmt.Printf("fib(%d)        = %d\n", n, fib(n))
	fibN := fibDynamic(n)
	fmt.Printf("fibDynamic(%d) = %v\n", n, fibN.String())
	// fibDynamic(1_000_000) // stackoverflow
}

func fib(n uint) uint {
	if n == 0 || n == 1 {
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
