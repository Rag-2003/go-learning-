package main

import "fmt"

func simpleFunction() {
	fmt.Println("heelo function")
}
func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("i am learning function in  golang")

	simpleFunction()

	ans := add(4, 5)
	fmt.Println("add of two numbsers ", ans)
}
