package main

import (
	"fmt"
	"net"
)

func main() {
	// Ожидание запросов от клиентов
	message := "Hello, I am a server" // отправляемое сообщение
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Accepted request from: ", conn.RemoteAddr())
		conn.Write([]byte(message))
		conn.Close()
	}
}
