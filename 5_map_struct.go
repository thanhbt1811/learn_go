package main

import (
	"fmt"
)

/*
1. Map
2. Struct
*/
type Student struct {
	fullName string
	age      int
	isMale   bool
	subject  []string
}
type Address struct {
	city     string
	district string
	ward     string
}

func main() {
	fmt.Println("Map and Struct")
	//1. Map
	var studentNameAge = map[string]int{
		"BTT": 22,
		"NDT": 21,
		"LTN": 22,
	}
	fmt.Println(studentNameAge)
	citiesIdMap := make(map[string]int)
	citiesIdMap = map[string]int{
		"HN":  1,
		"DN":  2,
		"HCM": 3,
	}
	fmt.Println(citiesIdMap)
	fmt.Println(citiesIdMap["HN"])
	fmt.Println(citiesIdMap["HCM"])
	citiesIdMap["HP"] = 4
	fmt.Println(citiesIdMap)
	delete(citiesIdMap, "HP")
	fmt.Println(citiesIdMap)
	_, isExist := studentNameAge["HP"]
	fmt.Println(isExist)
	//2. Struct
	var student = Student{
		fullName: "BTT",
		age:      22,
		isMale:   true,
		subject: []string{
			"Math", "English", "Computer",
		},
	}
	fmt.Println(student)
	fmt.Println(student.fullName)
	var address = struct {
		city     string
		district string
		ward     string
	}{
		city:     "HN",
		district: "TT",
		ward:     "TT",
	}
	fmt.Println(address)

}
