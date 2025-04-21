// package main

// import "fmt"

// func early() int {
// 	defer fmt.Println("the vlue shuld")
// 	return 1

// }
// func value() {
// 	for i := 0; i <= 3; i++ {
// 		defer fmt.Println(i)
// 	}
// }
// func main() {
// 	// defer new(1)
// 	// new(3)
// 	// defer new(4)
// 	// value()
// 	// early()

// }

//package main

// import (
//     "log"
//     "os"
// )

// func readFile() {
//     file, err := os.Open("temp.txt")
//     if err != nil {
//         log.Fatal("Error opening file:", err)
//     }
//     defer file.Close()

//     // Perform operations on the file
//     log.Println("File opened successfully")
// }

// func main() {
// 	 readFile()
// }

// return value as defer :  defer function executeafter return statement but beforethe function actully return 
package main

import "fmt"

func modifyReturn() (result int) {
	defer func() {
		result--
	}()
	return 0
}

func main() {
	fmt.Println(modifyReturn()) // Output: 1
}
