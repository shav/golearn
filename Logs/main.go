package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Открытие лог-файла
	logFile, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()

	fmt.Println("--------------------------------------")

	// Запись в лог-файл
	logger := log.New(logFile, "prefix: ", log.LstdFlags|log.Lshortfile)
	logger.Println("text to append")
	logger.Println("more text to append")

	fmt.Println("--------------------------------------")

	// Задизабливание логирования
	logger.SetOutput(ioutil.Discard)
	logger.Println("Hello, world!")
}
