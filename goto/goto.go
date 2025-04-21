package main

import "fmt"

func gotomain() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			if i == 1 && j == 1 {
				goto end
			}
			fmt.Println("dfgf",i,j)
		}
	}

end:
fmt.Println("exit the loop")
}

func main() {
  gotomain()
}
