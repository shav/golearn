package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Соединение с сервером и отправка запросов
	httpRequest := "GET / HTTP/1.1\n" +
		"Host: golang.org\n\n"
	conn, err := net.Dial("tcp", "yandex.ru:80")
	if err != nil {
		fmt.Println(err)
	} else {
		defer conn.Close()

		if _, err = conn.Write([]byte(httpRequest)); err != nil {
			fmt.Println(err)
			return
		}

		io.Copy(os.Stdout, conn)
		fmt.Println("Done")
	}

	fmt.Println("----------------------------------------")

	conn, err = net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
	} else {
		defer conn.Close()
		io.Copy(os.Stdout, conn)
		fmt.Println("\nDone")
	}

	fmt.Println("----------------------------------------")

	conn, err = net.Dial("tcp", "127.0.0.1:4547")
	if err != nil {
		fmt.Println(err)
	} else {
		defer conn.Close()

		words := []string{"red", "green", "black"}
		for _, word := range words {
			// отправляем запрос на сервер
			n, err := conn.Write([]byte(word))
			if n == 0 || err != nil {
				fmt.Println(err)
				continue
			}
			// получаем ответ
			response := make([]byte, 1024)
			n, err = conn.Read(response)
			if err != nil {
				continue
			}
			translate := string(response[0:n])
			fmt.Printf("Перевод: %s - %s\n", word, translate)
		}
		fmt.Println("\nDone")
	}
}
