package main

import (
	"artem/functions/temperature"
	"fmt"
)

func main() {
	var c temperature.Celcius = +101
	fmt.Printf("c.ToFarenheite(): %0.1f\n", c.ToFarenheite())

	var f temperature.Farenheite = -10
	fmt.Printf("f.ToCelcius(): %0.1f\n", f.ToCelcius())
}
