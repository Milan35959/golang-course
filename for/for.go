package main

import "fmt"

//for-> only construct in go for looping
func main() {
	//while loop

	// i := 1
	// for i <= 3 {
	// 	fmt.Print(i)
	// 	i = i + 1
	// }

	//infinite loop
	// for {
	// 	println("1")
	// }

	//classic for loop

	// for i := 0; i <= 3; i++ {
	// 	// break
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)

	// }

	//range loop

	for i := range 10 {
		fmt.Println(i)
	}
}
