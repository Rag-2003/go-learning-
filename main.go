package main

import (
	"fmt"
	// "golang-learn1/myutil"
)

func main() {
	// fmt.Println("hello go lang")
	// myutil.PrintMessage("my name is")
	// myutil.PrintMessage("my age")
	// var name string = "Raghav"
	// fmt.Println(name)
	var money int = 43000
	fmt.Println(money)
	var currency = 3400
	fmt.Println("this is  my currency", currency)

	//float example
	var dimentions float64 = 87.34
	fmt.Println("dimentions are", dimentions)

	//boolean example
	var decidsed bool = false //In  boolean  data type we can  change the value from  true to  false  but in  integer or any other data type we can reassign  the variable accept boolean
	decidsed = true
	fmt.Println("decidsed is", decidsed)

	//one ore kind of variable we can defined as :

	 const person = "valueis raghav"
	 const person1 = 123
	fmt.Println(person,person1)

	

}
