package main

import "fmt"

func divide(a, b float64) (float64, error) { //in  this we are getting the main function  as argument where a,b will be int but outside the function  we are  having an onother value of  the new value
	if b == 0 {
		return 0, fmt.Errorf("denominator is not be 0 ") // its not necessary to  write the fmt.errorf we can  directly pass the string thst set.
	}
	return a / b, nil

}

func main() {
	fmt.Println("starting error handling in  the function ")
	ans, err := divide(10, 0)

	if err != nil {

		fmt.Println("error handlingn")
	}
	fmt.Println("division  of the 2 number is ", ans)
}


// package main

// import "fmt"

// func divide(a, b float64) (float64, string) { //in  this we are getting the main function  as argument where a,b will be int but outside the function  we are  having an onother value of  the new value
// 	if b == 0 {
// 		return 0, "he" // its not necessary to  write the fmt.errorf we can  directly pass the string thst set.
// 	}
// 	return a / b, "nil"

// }

// func main() {
// 	fmt.Println("starting error handling in  the function ")
// 	ans, err := divide(10, 0)

// 	if err != "nil" {

// 		fmt.Println("error handlingn")
// 	}
// 	fmt.Println("division  of the 2 number is ", ans)
// }
