package main

import "fmt"

func new(i int) {
	fmt.Println("the vlue shuld", i)

}
func main() {
	defer new(1)
	new(3)
	defer new(4)
}