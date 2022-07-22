package main

import (
	"fmt"
	"math"
	"math/big"
	"math/cmplx"
)

func main() {
	// Целые числа
	var int0 int
	var int320 int32
	var int640 int64
	var uint320 uint32
	var uint640 uint64

	int320 = 1 << 30
	int0 = int(int320) // повышающее преобразование
	fmt.Printf("int32 -> int: %d\n", int0)

	int640 = 1 << 31
	int320 = int32(int640) // понижающее преобразование
	fmt.Printf("int64 -> int32: %d\n", int320)

	int320 = 1 << 30
	int640 = int64(int320) // повышающее преобразование
	fmt.Printf("int32 -> int64: %d\n", int640)

	// int320 = -(1 << 30)
	// uint640 = int(int320) // несовместимое преобразование
	// fmt.Println(uint640)

	uint640 = 1 << 63
	int320 = int32(uint640) // понижающее преобразование
	fmt.Printf("uint -> int32: %d\n", int320)

	uint640 = 1 << 32
	uint320 = uint32(uint640) // понижающее преобразование
	fmt.Printf("uint64 -> uint32: %d\n", uint320)

	int0 = 1
	int320 = 1
	fmt.Println(int0 == int(int320))
	fmt.Println(int0 < int(int320))

	sum := int0 + int(int320)
	fmt.Printf("sum = %d \n", sum)

	int640 = 1 << 62
	var int641 int64 = 1 << 62
	sumOverflow := int640 + int641
	fmt.Printf("sumOverflow = %d \n", sumOverflow)
	multiOverflow := int640 * int641
	fmt.Printf("multiOverflow = %d \n", multiOverflow)

	var int642 int64 = 2.5e+18
	fmt.Println(int642)

	d1 := 3
	d2 := 2
	fmt.Printf("+int / +int: %[1]d = %[3]d * %[2]d %+[4]d\n", d1, d2, d1/d2, d1%d2)
	fmt.Printf("-int / -int: %[1]d = %[3]d * %[2]d %+[4]d\n", -d1, -d2, -d1/-d2, -d1%-d2)
	fmt.Printf("-int / +int: %[1]d = %[3]d * %[2]d %+[4]d\n", -d1, d2, -d1/d2, -d1%d2)
	fmt.Printf("+int / -int: %[1]d = %[3]d * %[2]d %+[4]d\n", d1, -d2, d1/-d2, d1%-d2)

	//++d1 // prefix increment is not supported
	//var intInc = (d1++) + 1 // increment in expression is not supported
	d1++
	fmt.Printf("int++: %d\n", d1)
	d1--
	fmt.Printf("int--: %d\n", d1)
	d1 += 10
	fmt.Printf("int+=: %d\n", d1)

	// // Error:
	// var zero int = 0
	// fmt.Printf("div by zero: %d\n", 1/zero)

	fmt.Println("--------------------------------------")

	// Вещественные числа
	var f0 float32 = 3.14
	var f1 float32 = 3.1401
	var f2 float32 = 3.5001
	fmt.Println(f0 == f1)
	fmt.Println(f0 < f1)
	fmt.Println(f0 + f1)
	fmt.Println(f0 * f1)
	fmt.Println(10.0 / 4)
	fmt.Println(10.0 / 3.14)
	fmt.Printf("float -> int: %d\n", int(f2))
	fmt.Printf("round(float): %d\n", int(math.Round(float64(f2))))

	var f3 float64 = 2.56789e+100
	fmt.Println(f3)

	fmt.Printf("int / float: %f\n", float32(int320)/f1)

	f0++
	fmt.Printf("float++: %f\n", f0)
	f0--
	fmt.Printf("float--: %f\n", f0)

	fmt.Println("--------------------------------------")

	// Комплексные числа
	var c0 complex64 = 1 + 2i
	var c1 complex64 = 1 + 2i
	fmt.Println(c0 == c1)
	//fmt.Println(c0 < c1) // unsupported
	fmt.Println(c0 + c1)
	fmt.Println(c0 * c1)
	fmt.Println(c0 / c1)

	var c3 complex64 = 2.3456e+10 + 3.14e+20i
	fmt.Println(c3)

	var c4 = complex(1, 2)
	fmt.Println(c4)

	c0++
	fmt.Printf("complex++: %v\n", c0)
	c0--
	fmt.Printf("complex--: %v\n", c0)

	fmt.Println("--------------------------------------")

	// Math
	fmt.Printf("abs: %d\n", int(math.Abs(-10)))
	fmt.Printf("sin: %f\n", math.Sin(math.Pi/4))
	fmt.Printf("log10: %f\n", math.Log10(1000))
	fmt.Printf("exp: %f\n", math.Exp(2))
	fmt.Printf("sqrt: %f\n", math.Sqrt(144))
	fmt.Printf("inf * -inf: %f\n", math.Inf(+1)*math.Inf(-1))

	fmt.Printf("exp complex: %v\n", cmplx.Exp(2+3i))
	fmt.Printf("pow complex: %v\n", cmplx.Pow(2+3i, -1i))
	ro, phi := cmplx.Polar(1 - 2i)
	fmt.Printf("polar complex: %f*e^(i*%f)\n", ro, phi)
	fmt.Printf("complex from polar: %v\n", cmplx.Rect(ro, phi))

	fmt.Println("--------------------------------------")

	// Big numbers
	var n1 = big.NewInt(1e+10)
	fmt.Println(n1)
	var n2, _ = new(big.Int).SetString("1234567890098765432100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 10)
	fmt.Println(n2)
	var n3 = new(big.Int).SetBytes([]byte{1, 2}) // в 256-ричной системе счисления
	fmt.Println(n3)
	fmt.Println("big int multi: ", new(big.Int).Mul(n1, n2))

	var q0 = big.NewRat(1, 2)
	fmt.Println(q0)
	var q1 = big.NewRat(4, 6) // дробь сразу автоматически сокращается
	fmt.Println(q1)
	var q2 = new(big.Rat).SetFloat64(0.4)
	fmt.Println(q2)
	var q3, _ = new(big.Rat).SetString("5/6")
	fmt.Println(q3)
	fmt.Println("big rational multi: ", new(big.Rat).Mul(q0, q1))

	var r0 = big.NewFloat(3.14485690238903427689034772348907199974235780923457890234759043625907823901457890234).SetPrec(100)
	fmt.Println(r0)
	var r1 = new(big.Float).SetRat(big.NewRat(2, 3)).SetPrec(100) // Что-то пошло не так
	fmt.Println(r1)
	fmt.Println("big float multi: ", new(big.Float).Mul(r0, r1).SetPrec(1000))
}
