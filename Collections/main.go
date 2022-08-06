package main

import (
	"artem/collections/mycollections"
	"fmt"
)

type FlagsType struct {
	F0 mycollections.Flag
	F1 mycollections.Flag
	F2 mycollections.Flag
}

var Flags = FlagsType{
	F0: 1 << 0,
	F1: 1 << 1,
	F2: 1 << 2,
}

func main() {
	bits := mycollections.BitSet(0)
	bits.Set(Flags.F2 | Flags.F1)
	fmt.Printf("%b\n", bits)
	bits.Toggle(Flags.F1)
	fmt.Printf("%b\n", bits)
	bits.Clear(Flags.F2)
	fmt.Printf("%b\n", bits)

	fmt.Println("--------------------------------------")
}
