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
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")

	fmt.Println("----------------------------------------")
}
