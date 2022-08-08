package main

import (
	"embed"
	"fmt"
)

// Считывает содержимое текстового файла
//go:embed MyFolder/readme.txt
var fileString string

// Считывает содержимое файла (в байтах)
//go:embed MyFolder/readme.txt
var fileBytes []byte

//go:embed MyFolder/readme.txt
//go:embed MyFolder/*.hash
var folder embed.FS

func main() {
	fmt.Println(fileString)
	fmt.Println(string(fileBytes))

	content1, err := folder.ReadFile("MyFolder/file1.hash")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content1))
	}

	content2, err := folder.ReadFile("MyFolder/bad.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content2))
	}
}
