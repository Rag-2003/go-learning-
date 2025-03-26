package main

import "fmt"

func main() {
	fmt.Println("I'm learning array in  golang")
	var name [5]string
	// 0 1 2 3 4
	name[0] = "prince"
	name[2] = "raghav"

	// fmt.Println("print the name of the values", name)
	// var number = [5]int{1, 2, 3, 4, 5}
	// fmt.Println("the number are : ", number)

	// fmt.Println("length of number of array is : ", len(number))
	fmt.Println("value of name at 2nd index",name[2])
	 
	var rag [10]string

	fmt.Println("enter the value in the array", rag)
}
