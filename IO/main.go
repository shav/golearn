package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type PhoneReader string

func (phone PhoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(phone); i++ {
		if phone[i] >= '0' && phone[i] <= '9' {
			p[count] = phone[i]
			count++
		}
	}
	return count, io.EOF
}

type PhoneWriter struct{}

func (writer PhoneWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Print(string(bs[i]))
		}
	}
	fmt.Println()
	return len(bs), nil
}

type Person struct {
	Name string
	Age  int
}

func main() {
	// Кастомный reader
	phone1 := PhoneReader("+1(234)567 9010")
	phone2 := PhoneReader("+2-345-678-12-35")

	buffer := make([]byte, len(phone1))
	phone1.Read(buffer)
	fmt.Println(string(buffer))

	buffer = make([]byte, len(phone2))
	phone2.Read(buffer)
	fmt.Println(string(buffer))

	fmt.Println("--------------------------------------")

	// Кастомный writer
	bytes1 := []byte("+1(234)567 9010")
	bytes2 := []byte("+2-345-678-12-35")

	pw := PhoneWriter{}
	pw.Write(bytes1)
	pw.Write(bytes2)

	fmt.Println("--------------------------------------")

	// открытие файла для чтения
	f1, err := os.OpenFile("hello.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	f1.Close()

	if exists, _ := FileExists("hello.txt"); !exists {
		// Создание файла
		f3, err := os.Create("hello.txt")
		if err != nil {
			fmt.Println(err)
		}
		f3.Close()
	}

	// открытие файла для записи
	f2, err := os.OpenFile("hello.txt", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	f2.Close()

	fmt.Println("--------------------------------------")

	// Запись в файл
	wFile, err := os.OpenFile("hello.txt", os.O_WRONLY, 0777)
	defer wFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		wFile.WriteString("Hello, World!\n")

		var person = Person{Name: "Artem", Age: 20}
		data, _ := json.Marshal(person)
		wFile.Write(data)
	}

	fmt.Println("--------------------------------------")

	// Чтение из файла
	rFile, err := os.Open("hello.txt")
	defer rFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		data := make([]byte, 64)
		for {
			n, err := rFile.Read(data)
			if err == io.EOF { // если конец файла
				break
			}
			fmt.Print(string(data[:n]))
		}
	}
}

func FileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
