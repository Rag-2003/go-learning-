package main

import "fmt"

func incr(n *int) int {
	return *n + 1
}
func main() {

	x := 40
	p := &x

	fmt.Printf("address of x: %d\n", p)
	fmt.Printf("address of x: %d\n", *p)

	*p = incr(p) //incr(&x)
	fmt.Printf("vlaue of x: %d\n", x)
}

