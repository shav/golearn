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

	fmt.Println("--------------------------------------")

	// Строки
	var emptyStr string
	fmt.Println(emptyStr)

	str0 := "peace0"
	fmt.Println(str0)

	var str1 = "peace1"
	fmt.Println(str1)

	var str2 string = "peace2"
	fmt.Println(str2)

	var unicodeStr string = "abcπγ本"
	fmt.Println(unicodeStr)

	var multilineStr1 = "line1\nline2"
	fmt.Println(multilineStr1)

	var multilineStr2 = `line3
line4
line5`
	fmt.Println(multilineStr2)

	var rawStr = `line1\nline2`
	fmt.Println(rawStr)

	fmt.Printf("ascii str[1]: %c\n", str0[1])
	fmt.Printf("unicode str[1]: %c\n", unicodeStr[1])
	fmt.Printf("unicode str[3:5]: %s\n", unicodeStr[3:5])
	fmt.Printf("unicode str[7:10]: %s\n", unicodeStr[7:10])
	//fmt.Println(emptyStr[1]) // index out of range

	//unicodeStr[0] = 'd' // строки являются неизменяемыми

	fmt.Printf("ascii str len: %d bytes, %d chars\n", len(str0), utf8.RuneCountInString(str0))
	fmt.Printf("unicode str len: %d bytes, %d chars\n", len(unicodeStr), utf8.RuneCountInString(unicodeStr))

	for _, c := range str0 {
		fmt.Printf("%c ", c)
	}
	fmt.Println()

	for _, c := range unicodeStr {
		fmt.Printf("%c ", c)
	}
	fmt.Println()
}
