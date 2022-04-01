package main

import (
	"fmt"
	"strconv"
)

//Declare factorial function here

func factorial (n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	//Insert Code here
	var userInput string
	fmt.Println("Enter a number: ")
	fmt.Scanln(&userInput)

	value, _ := strconv.ParseInt(userInput, 10, 0)

	fmt.Println(factorial(int(value)))

}
