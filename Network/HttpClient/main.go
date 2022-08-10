package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// Запрос через http-клиент
	client := http.Client{
		Timeout: 6 * time.Second,
	}
	response, err := client.Get("http://golang.org")
	if err != nil {
		fmt.Println(err)
	} else {
		defer response.Body.Close()
		io.Copy(os.Stdout, response.Body)
	}

	fmt.Println("\n-------------------------------------------------------------------------------------\n")

	client2 := &http.Client{}
	req, err := http.NewRequest(
		"GET", "https://google.com", nil,
	)
	req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent

	response, err = client2.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		defer response.Body.Close()
		io.Copy(os.Stdout, response.Body)
	}
}
