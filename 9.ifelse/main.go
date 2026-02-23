package main

import "fmt"

func main() {
	fmt.Println("we are going to learn if-else condition")
	var inputNum int
	fmt.Println("Please enter a valid number")
	fmt.Scan(&inputNum)
	if inputNum == 0 {
		fmt.Println("Number is Zero")
	} else if inputNum > 0 {
		fmt.Println("Number is Positive")
	} else {
		fmt.Println("Number is Negative")
	}

}
