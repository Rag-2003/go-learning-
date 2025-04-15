package main

import (
	"fmt"
)

func main() {

	// for loop
	for i := 0; i < 5; i++ {

		fmt.Printf("%d\n", i)
	}
	// for as while loop
	x := 10
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// infinite loop
	x = 0
	for {
		if x > 10 {
			break
		}
		fmt.Println(x)
		x++
	}
	//imp  topic for learning and implementing 
	//range for loop
	fmt.Println("Range for loop..........")
	number := []int{1, 2, 3,4,5,6}
	for idx, value := range number {
		fmt.Println(idx, value)
	}

}
