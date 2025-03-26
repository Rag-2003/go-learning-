package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hey, what's your name?")
	// var name string
	// var val int

	// fmt.Scan(&name)

	// fmt.Println("Hello, Mr.", name)
	// fmt.Println("whats your salary??")
	// fmt.Scan(&val)

	// fmt.Println("your salary amount is credit Rs ", val)
	 reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	 fmt.Println("hello, mr.", name)
}
