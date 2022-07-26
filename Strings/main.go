package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
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
	beta := gamma - 1
	fmt.Printf("%c: %[1]v\n", beta)
	delta := gamma + 1
	fmt.Printf("%c: %[1]v\n", delta)

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

	fmt.Println("--------------------------------------")

	// Поиск подстрок
	// Вариант 1
	var refString = "Mary had a little lamb"
	lookFor := "lamb"
	contain := strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	lookFor = "wolf"
	contain = strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	// Вариант 2
	startsWith := "Mary"
	starts := strings.HasPrefix(refString, startsWith)
	fmt.Printf("The \"%s\" starts with \"%s\": %t \n", refString, startsWith, starts)

	endWith := "lamb"
	ends := strings.HasSuffix(refString, endWith)
	fmt.Printf("The \"%s\" ends with \"%s\": %t \n", refString, endWith, ends)

	// Вариант 3
	refString = `[{ \"email\": \"email@example.com\" \"phone\": 555467890},
{ \"email\": \"other@domain.com\" \"phone\": 555467890}]`
	emailRegexp := regexp.MustCompile("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}")
	email1 := emailRegexp.FindString(refString)
	fmt.Println("First email: ", email1)

	allEmails := emailRegexp.FindAllString(refString, -1)
	fmt.Println("All emails: ")
	for _, email := range allEmails {
		fmt.Println(email)
	}

	fmt.Println("--------------------------------------")

	// разбиение на подстроки
	// Вариант 1
	refString = "Mary had a little lamb"
	words := strings.Fields(refString)
	fmt.Println("Split into words:")
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
	fmt.Println("--------------------------------------")

	// Вариант 2
	refString = "Mary_had a little_lamb"
	words = strings.Split(refString, "_")
	fmt.Println("Split into words by _:")
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
	fmt.Println("--------------------------------------")

	// Вариант 3
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_", r)
	}

	refString = "Mary*had,a%little_lamb"
	words = strings.FieldsFunc(refString, splitFunc)
	fmt.Println("Split into words by punctuation:")
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
	fmt.Println("--------------------------------------")

	// Вариант 4
	refString = "Mary*had,a%little_lamb"
	words = regexp.MustCompile("[*,%_]{1}").Split(refString, -1)
	fmt.Println("Split into words by punctuation via regexp:")
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
	fmt.Println("--------------------------------------")

	// объединение массива строк
	var refStringSlice = []string{
		"FIRST_NAME = 'Jack'",
		"INSURANCE_NO = 333444555",
		"EFFECTIVE_FROM = SYSDATE"}
	sentence := strings.Join(refStringSlice, " AND ")
	fmt.Println(sentence)

	fmt.Println("--------------------------------------")

	// конкатенация строк
	// Вариант 1
	var concatStr = str0 + unicodeStr
	fmt.Println(concatStr)

	myStrings := []string{"This ", "is ", "even ", "more ", "performant "}
	// Вариант 2
	buffer := bytes.Buffer{}
	for _, str := range myStrings {
		buffer.WriteString(str)
	}
	fmt.Println(buffer.String())

	// Вариант 3
	var strlen = 0
	for _, str := range myStrings {
		strlen += len(str)
	}
	concat := make([]byte, strlen)
	index := 0
	for _, str := range myStrings {
		index += copy(concat[index:], []byte(str))
	}
	fmt.Println(string(concat[:]))

	// Вариант 4
	var sb strings.Builder
	for _, str := range myStrings {
		sb.WriteString(str)
	}
	fmt.Println(sb.String())

	fmt.Println("--------------------------------------")

	// Замена подстрок
	// Вариант 1
	refString = "Mary had a little lamb"
	refString2 := "lamb lamb lamb lamb"

	out := strings.Replace(refString, "lamb", "wolf", -1)
	fmt.Println(out)
	out = strings.Replace(refString2, "lamb", "wolf", 2)
	fmt.Println(out)

	// Вариант 2
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")
	out = replacer.Replace(refString)
	fmt.Println(out)

	// Вариант 3
	regex := regexp.MustCompile("l[a-z]+")
	out = regex.ReplaceAllString(refString, "replacement")
	fmt.Println(out)
}
