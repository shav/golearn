package main

import (
	"fmt"
	"net"
)

var translate = map[string]string{
	"red":    "красный",
	"green":  "зеленый",
	"blue":   "синий",
	"yellow": "желтый",
}

func main() {
	listener, err := net.Listen("tcp", ":4547")

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
			conn.Close()
			continue
		}
		go handleConnection(conn) // запускаем горутину для обработки запроса
	}
}

// обработка подключения
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// считываем полученные в запросе данные
		// TODO: Если размер запроса превышает размер буфера, то нужно считывать данные запроса в цикле
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		source := string(input[0:n])
		// на основании полученных данных получаем из словаря перевод
		target, ok := translate[source]
		if ok == false { // если данные не найдены в словаре
			target = "unknown"
		}
		// выводим на консоль сервера диагностическую информацию
		fmt.Println(source, "-", target)
		// отправляем данные клиенту
		conn.Write([]byte(target))
	}
}
