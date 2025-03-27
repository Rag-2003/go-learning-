package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6,7,8,9,10)
	fmt.Println("number has data type", numbers)
	fmt.Printf("How to check  the type %T\n", numbers)
	fmt.Println("length:", len(numbers))
}
func another(){
   main := 123
   fmt.Println("the number should be more modified in  this format to ",  main)
}
