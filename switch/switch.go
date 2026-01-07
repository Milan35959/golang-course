package main

import (
	"fmt"
)

func main() {
	// simple switch

	// 	i := 6

	// 	switch i {
	// 	case 1:
	// 		fmt.Println("one")
	// 	case 2:
	// 		fmt.Println("two")
	// 	case 3:
	// 		fmt.Println("three")
	// 	case 4:
	// 		fmt.Println("four")
	// 	case 5:
	// 		fmt.Println("five")
	// 		// default:
	// 		// 	fmt.Println("not found")
	// 	}
	// }

	//multiple case switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's weekend")
	// default:
	// 	fmt.Println("It's a workay")
	// }

	//type switch

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I'm an integer")
		case string:
			fmt.Println("I'm a string")
		case bool:
			fmt.Println("I'm a boolean")
		default:
			fmt.Println("others")
		}
	}

	whoAmI(1)
}
