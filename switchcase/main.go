package main

import (
	"fmt"
)

func main() {

	// }
	day := 2
	switch day {
	case 1:
		fmt.Println("it's on monday")
	case 2:
		fmt.Println("its on tuesday")
	default:
		fmt.Println("df")
	}

	switch day {
	case 1, 2, 3:
		fmt.Println("cvdf")
	case 4, 5:
		fmt.Println("cvdv")
	default:
		fmt.Println("fSFF")
	}

	var x interface{} = 6
	switch x.(type) {
	case int:
		fmt.Println("x is integer")
	}
}
