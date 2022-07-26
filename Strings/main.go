package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Символы
	aChar0 := 'A'
	var aChar1 = 'A'
	var aChar2 rune = 'A'
	var star byte = '*'
	fmt.Printf("%c%c%c%c\n", aChar0, aChar1, aChar2, star)

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)

	var str = string([]rune{pi, alpha, omega, rune(bang)})
	fmt.Println(str)

	var piStr = string(pi)
	fmt.Println(piStr)

	const gammaStr = "γ"
	var gamma, _ = utf8.DecodeRuneInString(gammaStr)
	fmt.Printf("%c: %[1]v\n", gamma)

	const rootStr = "本"
	var root, _ = utf8.DecodeRuneInString(rootStr)
	fmt.Printf("%c: %[1]v (%d bytes)\n", root, len(rootStr))
	fmt.Println([]byte(rootStr))
}
