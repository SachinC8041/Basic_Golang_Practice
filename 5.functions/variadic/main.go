package main

import "fmt"

func main() {
	fmt.Println("We are going to learn Variadic Functions")

	//variadic function
	fmt.Println("Sum of first 4 Numbers is", findSumOfInts(10, 20, 30, 40, 50))

	//passing slice to variadic function
	nums := []int{10, 20, 30, 40}
	fmt.Println("The addition for the slice nums is:", findSumOfInts(nums...))

	//Defer Keyword practice

}

//Variadic Function
func findSumOfInts(params ...int) int {
	var sum int = 0
	for _, v := range params {
		sum += v
	}
	return sum
}
