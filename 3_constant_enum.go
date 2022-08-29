package main

import "fmt"

/*
1. Hằng số
2. Enum
3. Tự động khởi tạo giá trị cho enum iota
*/
const (
	max int16 = 365
	min       = 1
)

func main() {
	fmt.Println("Constants and Enum")
	//1. Hằng số
	const max = 30
	fmt.Printf("%v %T\n", max, max)
	fmt.Printf("%v %T\n", min, min)
	//2. Enum -> Multi constants
	const (
		red    = 0
		yellow = 1
		green  = 2
	)
	fmt.Printf("%v, %T\n", red, red)
	fmt.Printf("%v, %T\n", yellow, yellow)
	fmt.Printf("%v, %T\n", green, green)
	//3. iota
	const (
		init = (iota + 1) * 5
		loading
		success
		failed
	)
	println()
	fmt.Printf("%v, %T\n", init, init)
	fmt.Printf("%v, %T\n", loading, loading)
	fmt.Printf("%v, %T\n", success, success)
	fmt.Printf("%v, %T\n", failed, failed)
	println()
	const (
		_ = iota
		top
		bottom
		left
		right
	)
	fmt.Printf("%v, %T\n", top, top)
	fmt.Printf("%v, %T\n", bottom, bottom)
	fmt.Printf("%v, %T\n", left, left)
	fmt.Printf("%v, %T\n", right, right)
}
