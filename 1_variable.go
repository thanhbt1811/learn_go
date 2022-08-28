package main

import "fmt"
import "strconv"

var (
	name     = "Hello world"
	fullName = "Bui Tien Thanh"
)            // global variable
var n = 10   //Các package khác không thể truy cập
var M = 2000 //Các package khác có thể truy cập được
func main() {
	fmt.Println("Go Variable")
	//1. Khai báo biến
	var i int //Khai báo biến với kiểu dữ liệu không gán dữ liệu
	i = 10
	fmt.Printf("Gia tri %v, Kieu %T\n", i, i)
	var j = 20 // Khai báo biến với kiểu dữ liệu và dữ liệu
	fmt.Printf("Gia tri %v, Kieu %T\n", j, j)
	k := 30 // Khai báo biến ngầm hiểu kiểu dữ lệu
	fmt.Printf("Gia tri %v, Kieu %T\n", k, k)
	var l = "Hello" // Khai báo biến ngầm hiểu kiểu dữ lệu
	fmt.Printf("Gia tri %v, Kieu %T\n", l, l)

	//2. Global scope and bloc scope
	fmt.Printf("Gia tri %v, Kieu %T\n", name, name)
	fmt.Printf("Gia tri %v, Kieu %T\n", fullName, fullName)
	//3. Shadow
	/*
		Khi trong một hàm khai báo một biến cùng tên với biến toàn cục
		thì trong hàm đó sẽ sử dụng giá trị mới được khai báo
	*/
	fmt.Printf("Gia tri %v, Kieu %T\n", n, n)
	var n = 100
	fmt.Printf("Gia tri %v, Kieu %T\n", n, n)
	//4. Declared and not use : Không thể khai báo biến mà không sử dụng
	//5. Export global scope
	//6. Naming convention -> Camel
	//7. Convert type
	var n1 = 10
	var n2 = float32(n1)
	fmt.Printf("Gia tri cua n2 %v, Kieu %T\n", n2, n2)
	//Convert any type to String
	var stringN1 = strconv.Itoa(n1)
	fmt.Printf("Gia tri cua n2 %v, Kieu %T\n", stringN1, stringN1)
}
