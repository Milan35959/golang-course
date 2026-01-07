package main

import (
	"fmt"
	"maps"
)

// maps-> hash, object, dictionary
func main() {
	// creating map

	// m := make(map[string]string)

	// //setting elements
	// m["name"] = "GoLang"
	// m["are"] = "Backend"

	//get an element

	// fmt.Println(m["name"], m["are"])

	//IMP: if key doesnot exists that returns zero value of that type

	// m := make(map[string]int)
	// m["age"] = 10
	// m["year"] = 2024
	// fmt.Println(m["phone"])\

	// fmt.Println(len(m))

	// delete(m, "year")

	// clear(m)

	// fmt.Println(m)
	// fmt.Println(m)

	// m := map[string]int{"price": 40, "phones": 3}

	// k, ok := m["phones"]
	// fmt.Println(k)
	// if ok {
	// 	fmt.Println("price is", k)
	// } else {
	// 	fmt.Println("price key not found")
	// }

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 8}

	fmt.Println(maps.Equal(m1, m2))
}
