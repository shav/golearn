package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// Конкатенация путей
	file := filepath.Join("dir1", "dir2", "filename.png")
	fmt.Println("file: ", file)

	file2 := filepath.Join("dir1//", "filename.png")
	fmt.Println("file2: ", file2)

	file3 := filepath.Join("dir1/../dir2", "filename.png")
	fmt.Println("file3: ", file3)

	fmt.Println("\n--------------------------------------")

	// Составные части пути
	fmt.Println("Dir(f): ", filepath.Dir(file))

	filename := filepath.Base(file)
	fmt.Println("Filename(f): ", filename)

	ext := filepath.Ext(file)
	fmt.Println("Ext(f): ", ext)

	filenameWithoutExt := strings.TrimSuffix(filename, ext)
	fmt.Println("Name(f): ", filenameWithoutExt)

	fmt.Println("\n--------------------------------------")

	// Построение относительных путей
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rel)
	}

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rel)
	}
}
