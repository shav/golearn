package main

import (
	"fmt"
	"os"
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

	fmt.Println("\n--------------------------------------")

	// Создание файлов и папок
	err = os.Mkdir("subdir", 0755)
	if err != nil {
		fmt.Println(err)
	}
	//defer os.RemoveAll("subdir")

	createEmptyFile("subdir/file1.txt")
	os.MkdirAll("subdir/parent/child", 0755)
	createEmptyFile("subdir/parent/file2.txt")
	createEmptyFile("subdir/parent/file3.txt")
	createEmptyFile("subdir/parent/child/file4.txt")

	fmt.Println("\n--------------------------------------")

	// Чтение файловой структуры
	var folderName = "subdir/parent"
	folder, err := os.ReadDir(folderName)
	fmt.Printf("Listing %s:\n", folderName)
	// Перечисляет содержимое папки нерекурсивно (т.е. без содержимого подпапок)!
	for _, entry := range folder {
		fmt.Printf("    %s: %t\n", entry.Name(), entry.IsDir())
	}
	fmt.Println()

	folderName = "subdir/parent/child"
	os.Chdir(folderName)
	folder, _ = os.ReadDir(".")
	fmt.Printf("Listing %s:\n", folderName)
	for _, entry := range folder {
		fmt.Printf("    %s: %t\n", entry.Name(), entry.IsDir())
	}
	fmt.Println()

	// Рекурсивный обход папки
	folderName = "subdir"
	os.Chdir("../../..")
	fmt.Printf("Visiting %s:\n", folderName)
	filepath.Walk(folderName, visit)

	fmt.Println("\n--------------------------------------")

	// Временные файлы и папки
	tmpFile, err := os.CreateTemp("", "go-sample-*.txt")
	// defer os.Remove(tmpFile.Name())
	fmt.Println("Temp file name: ", tmpFile.Name())

	tmpFolder, err := os.MkdirTemp("", "sampledir")
	// defer os.RemoveAll(tmpFolder)
	fmt.Println("Temp dir name: ", tmpFolder)
}

func visit(entry string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("    %s: %t\n", entry, info.IsDir())
	return nil
}

func createEmptyFile(name string) {
	d := []byte("")
	err := os.WriteFile(name, d, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
