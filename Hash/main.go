package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

type HashAlgorithm interface {
	io.Writer
	Sum(in []byte) []byte
}

var fileName string

func init() {
	flag.StringVar(&fileName, "file", "", "file for calculating hash")
	flag.Parse()
}

func MD5(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func SHA1(data string) string {
	hash := sha1.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func SHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func SHA512(data string) string {
	hash := sha512.Sum512([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func FileHash(filePath string, hash HashAlgorithm) string {
	md5.New()
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
	// хеш строки
	str := "Hello"
	fmt.Println("str md5-hash: ", MD5(str))
	fmt.Println("str sha1-hash: ", SHA1(str))
	fmt.Println("str sha256-hash: ", SHA256(str))
	fmt.Println("str sha512-hash: ", SHA512(str))

	fmt.Println("--------------------------------------")

	// хеш файла
	if fileName == "" {
		fmt.Println("File name is not specified")
		return
	}
	fmt.Println("file md5-hash: ", FileHash(fileName, md5.New()))
	fmt.Println("file sha1-hash: ", FileHash(fileName, sha1.New()))
	fmt.Println("file sha256-hash: ", FileHash(fileName, sha256.New()))
	fmt.Println("file sha512-hash: ", FileHash(fileName, sha512.New()))
}
