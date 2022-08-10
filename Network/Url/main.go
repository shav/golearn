package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	urlStr := "https://admin:11111@myhost.com:5432/path/to/resource?param=value#/card/1234"
	url, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("scheme: ", url.Scheme)
		fmt.Println("user: ", url.User.Username())

		password, _ := url.User.Password()
		fmt.Println("password: ", password)

		fmt.Println(url.Host)
		host, port, _ := net.SplitHostPort(url.Host)
		fmt.Println("host: ", host)
		fmt.Println("port: ", port)

		fmt.Println("path: ", url.Path)
		fmt.Println("fragment: ", url.Fragment)

		fmt.Println("query: ", url.RawQuery)
		query := url.Query()
		fmt.Println(query)
	}
}
