package main

import (
	"fmt"
	"strconv"
)

// a := 12
// b := 13
// var c int

// c = a + b
// fmt.Println("the value should be ", c)

// c = a - b
// fmt.Println("the sub values us", c)

//   d := float64(a) / float64(b)
// fmt.Println("the division  value must be equal  to  main", d)

func Add(numbers ...float64) float64 {
	total := 0.0
	for i := range numbers {
		total += float64(numbers[i])

	}
	return total
}
func Convert(str string) (int, error) {
	return strconv.Atoi(str)
}

func main() {
	i, e := Convert("123")
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(i)

	i, e = Convert("gfdscvd")
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(i)
}
