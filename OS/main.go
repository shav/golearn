package main

import (
	"fmt"
	"os"
	"path/filepath"
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
}
