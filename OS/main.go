package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Получение аргументов командной строки
	if len(os.Args) != 3 {
		var programName = filepath.Base(os.Args[0])
		fmt.Println("Usage:", programName, "PATTERN", "FILE")
		return
	}
	pattern := os.Args[1]
	file := os.Args[2]
	fmt.Printf("%s: %s\n", file, pattern)

	fmt.Println("--------------------------------------")

	// Список переменных окружения
	for _, s := range os.Environ() {
		kv := strings.SplitN(s, "=", 2) // unpacks "key=value"
		fmt.Printf("%q = %q\n", kv[0], kv[1])
	}

	fmt.Println("--------------------------------------")

	// Модификация переменных окружения
	fmt.Printf("%q\n", os.Getenv("ARTEM"))
	os.Setenv("ARTEM", "/bin/dash")
	fmt.Printf("%q\n", os.Getenv("ARTEM"))
	os.Unsetenv("ARTEM")
	fmt.Printf("%q\n", os.Getenv("ARTEM"))
}
