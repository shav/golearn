package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
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
			response := make([]byte, 0)
			conn.SetReadDeadline(time.Now().Add(time.Second * 5))
			for {
				buffer := make([]byte, 1024)
				n, err = conn.Read(buffer)
				if err != nil {
					break
				}
				response = append(response, buffer[0:n]...)
				conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
			}
			translate := string(response[:])
			fmt.Printf("Перевод: %s - %s\n", word, translate)
		}
		fmt.Println("\nDone")
	}
}
