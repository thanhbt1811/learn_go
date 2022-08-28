package main

import "fmt"

/*
Primitive data type
1. Boolean
2. Numerics
- Integer
- Float
- Complex
3. Text
*/
func main() {
	fmt.Println("Primitive data type")
	//1. Boolean
	var a = true
	fmt.Printf("%v %T\n", a, a)
	//default value of boolean data type
	var b bool
	fmt.Printf("%v %T\n", b, b)
	//2. Numeric
	//2.1. Integer
	/*
		- signed integer - Số nguyên có dâu (int8, int32, int64)
		- unsigned integer - Số nguyên không dấu (uint8,uint32,uint64)
	*/
	//signed integer
	var c int8 = -23
	fmt.Printf("%v %T\n", c, c)
	//unsigned integer
	var d uint8 = 10
	var e uint8 = 5
	fmt.Printf("%v %T\n", d, d)
	fmt.Printf("%v + %v = %v\n", d, e, d+e)
	fmt.Printf("%v - %v = %v\n", d, e, d-e)
	fmt.Printf("%v * %v = %v\n", d, e, d*e)
	fmt.Printf("%v / %v = %v\n", d, e, d/e)
	fmt.Printf("%v / %v dư %v\n", d, e, d%e)
	/*
		Phép toán trên bit
		- AND: &
		- OR: |
		- NOT: ^
		- XOR: &^
	*/
	fmt.Printf("%v & %v = %v\n", d, e, d&e)
	fmt.Printf("%v | %v = %v\n", d, e, d|e)
	fmt.Printf("%v ^ %v = %v\n", d, e, d^e)
	fmt.Printf("%v &^ %v = %v\n", d, e, d&^e)
	/*
		Toàn tử dịch bit : <<, >>
	*/
	fmt.Printf("%v >> 3\n = %v", d, d>>3)
	fmt.Printf("%v << 3\n = %v", d, d<<3)
	//2.2. Float (float32,float64)
	var f = 3.14
	var g = 7.21
	fmt.Printf("%v + %v = %v\n", f, g, f+g)
	fmt.Printf("%v - %v = %v\n", f, g, f-g)
	fmt.Printf("%v * %v = %v\n", f, g, f*g)
	fmt.Printf("%v / %v = %v\n", f, g, f/g)

	//2.4. Complex
	var h complex64 = 1 + 3i
	fmt.Println(h)
	fmt.Printf("Thuc: %v\n", real(h))
	fmt.Printf("Ao: %v\n", imag(h))
	var i = 323 + 443i
	fmt.Println(i)
	//3. Text
	//3.1. String
	var j = "Hello "
	fmt.Println(j)
	var k = "Golang"
	fmt.Println(k)
	fmt.Println(j + k)
	//3.2. byte -> utf-8
	//3.3. Rune -> utf-32
}
