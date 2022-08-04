package main

import (
	crypto "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

type Guid string

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crypto.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func main() {
	// Псвевдослучайные числа (небезопасные)
	rand.Seed(time.Now().UnixNano())
	fmt.Println("random int: ", rand.Intn(10))
	fmt.Println("random int: ", rand.Int63())
	fmt.Println("random int: ", RandomInt(100, 200))

	fmt.Println("random float: ", rand.Float32())
	fmt.Println("random float: ", RandomFloat(10, 20))

	fmt.Printf("random char: %c\n", RandomChar())
	fmt.Printf("random string: %s\n", RandomString(10))
	fmt.Printf("random password: %s\n", RandomPassword(10))

	fmt.Printf("guid: %s\n", NewGuid())

	// Другая псевдослучайная последовательность чисел
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("random int: ", generator.Intn(10))

	fmt.Println("--------------------------------------")

	// Диапазон случайных целых чисел
	fmt.Println("random int range: ", rand.Perm(10))

	// Случайная перестановка элементов списка
	list := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	fmt.Println(list)

	fmt.Println("--------------------------------------")

	// Криптографически безопасные случайные числа
	fmt.Println("crypto random int: ", NewCryptoRand(100))
	// другой способ
	var src cryptoSource
	rnd := rand.New(src)
	fmt.Println("crypto random int: ", rnd.Intn(100))
}

func NewCryptoRand(max int64) int64 {
	rnd, err := crypto.Int(crypto.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return rnd.Int64()
}

func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomChar() rune {
	return 'a' + rune(rand.Intn('z'-'a'+1))
}

func RandomString(length uint) string {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	var sb strings.Builder
	for i := uint(0); i < length; i++ {
		sb.WriteRune(chars[generator.Intn(len(chars))])
	}
	return sb.String()
}

func RandomPassword(length uint) string {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	buf := make([]byte, length)
	buf[0] = digits[generator.Intn(len(digits))]
	buf[1] = specials[generator.Intn(len(specials))]
	for i := uint(2); i < length; i++ {
		buf[i] = all[generator.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

func NewGuid() Guid {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	var str = fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return Guid(str)
}
