package main

import (
	"fmt"
	"github.com/yourbasic/bit"
	"math"
	"math/big"
)

func main() {
	fmt.Println(getPrimeNumbers(100))

	fmt.Println("--------------------------------------")

	fmt.Println(BytesToString(123456))
	fmt.Println(BytesToString(1e9))

	fmt.Println("--------------------------------------")

	var num, _ = new(big.Int).SetString("170141183460469231731687303715884105727", 10)
	fmt.Printf("%v is prime: %t", num, IsPrime(*num))
}

func IsPrime(num big.Int) bool {
	return num.ProbablyPrime(20)
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

func BytesToString(bytes int64) string {
	const unit = 1000
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	//div, exp := int64(unit), 0
	//for n := bytes / unit; n >= unit; n /= unit {
	//	div *= unit
	//	exp++
	//}

	var exp uint = uint(math.Log10(float64(bytes)) / math.Log10(unit))
	var div float64 = math.Pow(unit, float64(exp))

	return fmt.Sprintf("%.1f %cB",
		float64(bytes)/float64(div), "kMGTPE"[exp-1])
}
