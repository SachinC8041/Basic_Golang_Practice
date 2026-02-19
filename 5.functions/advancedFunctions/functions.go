package main

import "fmt"

func main() {

	//storing function in variable
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(5, 5))

	//passing function as an argument
	fmt.Println(operate(12, 12, addition))

	//returning them from other functions
	double := multiplier(2)
	fmt.Println(double(5))

	//multiple return value section in function
	multiplicationAns, additionAns := operation(12, 12)
	fmt.Printf("Multiplication ans %d and Addition ans is %d\n", multiplicationAns, additionAns)

	//Ignoring the results using Blank Identifier
	product, _ := operation(12, 12)
	fmt.Println("Multiplication result is, ", product)

	//handling error
	divisionAns, err := divide(12, 0)
	if err != nil {
		fmt.Println("Error is", err)
	} else {
		fmt.Println("Division ans is", divisionAns)
	}

}

func operate(a int, b int, addition func(int, int) int) int {
	return addition(a, b)
}
func addition(a, b int) int {
	return a + b
}

func multiplier(factor int) func(int) int {

	return func(x int) int {
		return x * factor
	}
}

func operation(a int, b int) (int, int) {
	return a * b, a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't divide by zero")
	} else {
		return a / b, nil
	}
}
