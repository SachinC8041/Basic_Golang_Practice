package main

import "fmt"

func main() {
	var a, b int = 10, 4

	addition := a + b
	fmt.Println("Addition is :", addition)

	substraction := a - b
	fmt.Println("Substraction is :", substraction)

	multiplication := a * b
	fmt.Println("Multiplication is :", multiplication)

	division := a / b
	fmt.Println("Division is :", division)

	var div float64 = 22 / 7
	fmt.Println(div)

	var div2 float64 = 22 / 7.0 // or 22.0/7 as well
	fmt.Println(div2)

	//Note : in Go , Int / Int will always give int as division answer . (truncated towards zero)
	// if we want floting point answer then at least one of the two numbers should be float type
	// note : the exection of the operation will start from right to left i.e. 7.0 to 22 so compiler doesnt understand
	// the ans is going to be int or float

	reminder := a % b
	fmt.Println("Reminder is :", reminder)
}
