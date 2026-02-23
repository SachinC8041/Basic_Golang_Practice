package main

import "fmt"

func main() {
	fmt.Println("Switch statement practice session")

	var inputNum int
	fmt.Scan(&inputNum)

	switch inputNum {
	case 1, 3, 5, 7:
		fmt.Println("Odd week day")
	case 2, 4, 6:
		fmt.Println("Even week day")
	default:
		fmt.Println("Invalid input")
	}
}
