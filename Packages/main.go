package main

import (
	"fmt"
	"math/big"
	"rsc.io/quote"

	// Импорт пакета с github-а
	"github.com/yourbasic/graph"

	// Пакет импортитрован, но не используется (нужно для инициализации данных в пакете):
	_ "image/png"

	// Статический импорт:
	. "math"

	// Конфликты имён:
	crypto "crypto/rand"
	rand "math/rand"
)

func main() {
	fmt.Println(quote.Glass())

	fmt.Println(rand.Intn(10))
	rnd, _ := crypto.Int(crypto.Reader, big.NewInt(10))
	fmt.Println(rnd)

	// Данные функции статически импортированы из пакета math
	one := Pow(Sin(Pi/4), 2) + Pow(Cos(Pi/4), 2)
	fmt.Println(one)

	g := graph.New(5)
	fmt.Println(g)
}
