package main

import (
	"fmt"
)

/*
1. Array
2. Slice
*/
func main() {
	fmt.Println("Array and Slice")
	//1. Array
	var names = [...]string{"Alex", "Yu", "!"}
	fmt.Printf("%v\n", names[0])
	fmt.Printf("%v\n", names[1])
	fmt.Printf("names length %v\n", len(names))
	var points [3]int // Mảng có giá trị mặc định là 0
	fmt.Printf("%v\n", points)
	points[0] = 4
	points[1] = 5
	points[2] = 10
	fmt.Printf("%v\n", points[0])
	fmt.Printf("%v\n", points[1])
	fmt.Printf("%v\n", points[2])
	fmt.Printf("points length %v\n", len(points))
	var cities [3]string //Mảng có giá trị mặc định là ""
	fmt.Printf("%v\n", cities)
	cities[1] = "HCM"
	fmt.Printf("%v\n", cities)
	var copyPoints = &points
	copyPoints[1] = 8
	fmt.Println(copyPoints)
	fmt.Println(points)
	//2. Slice
	var countries = []string{"VN", "UK", "US", "JP", "CN", "KR", "SP"}
	fmt.Println(countries)
	var countries1 = countries[:]
	var countries2 = countries[3:]
	var countries3 = countries[:5]
	var countries4 = countries[3:5]
	fmt.Println(countries1)
	fmt.Println(countries2)
	fmt.Println(countries3)
	fmt.Println(countries4)
	var categories = make([]string, 0)
	fmt.Println(categories)
	categories = append(categories, "Library")
	fmt.Println(categories)
	categories = append(categories, "Room", "Class")
	fmt.Println(categories)

}
