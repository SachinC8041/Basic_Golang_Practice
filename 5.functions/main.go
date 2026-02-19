package main

import "fmt"

func simplePrint() {
	fmt.Println("Hello Suraj")
}

func addition(a int, b int) int {
	return a + b
}

func someFunction(a, b int) (int, int) {
	return a * b, a + b
}

func main() {
	fmt.Println("We are learning functions in Go lang")
	simplePrint()

	additionResult := addition(12, 32)
	fmt.Println("Addition is", additionResult)

	//use of underscore operator
	mulitplication, _ := someFunction(12, 12)
	fmt.Println(mulitplication)

	//Concept of Anonymous Function
	//Create Anonymous function and give it a call immediately
	func() {
		fmt.Println("This is Anonymous Function and it is getting called immediately")
	}()

	//Create Anonymous function and give it a call later by assigning it to a variable and calling it later
	greeting := func() {
		fmt.Println("This is Anonymous function and we are assigning it to some variable")
	}
	greeting()

	//Read about First Class Citizens and First Class Objects in Functions in GO

	//In Go, Functions are First Class Citizens
	// 	That means:
	// ✔ You can store them in variables
	// ✔ Pass them as arguments
	// ✔ Return them from other functions
}
