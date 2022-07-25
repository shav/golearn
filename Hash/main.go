package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "file", "", "file for calculating hash")
	flag.Parse()
}

func MD5(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func FileMD5(filePath string) string {
	hash := md5.New()
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(hash, file)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func main() {
	// md5-хеш строки
	str := "Hello"
	fmt.Println("str md5-hash: ", MD5(str))

	// md5-хеш файла
	if fileName == "" {
		fmt.Println("File name is not specified")
		return
	}
	fmt.Println("file md5-hash: ", FileMD5(fileName))
}
