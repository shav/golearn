package main

import (
	"fmt"
	"github.com/yourbasic/bit"
	"math"
)

func main() {
	fmt.Println(getPrimeNumbers(100))

	fmt.Println("--------------------------------------")
}

func getPrimeNumbers(max int) []uint {
	// Sieve of Eratosthenes
	sieve := bit.New().AddRange(2, max)
	sqrtN := int(math.Sqrt(float64(max)))
	for p := 2; p <= sqrtN; p = sieve.Next(p) {
		for k := p * p; k < max; k += p {
			sieve.Delete(k)
		}
	}

	var result = make([]uint, 0, sieve.Size())
	for p := 2; p <= max && p > 0; p = sieve.Next(p) {
		result = append(result, uint(p))
	}

	return result
}
