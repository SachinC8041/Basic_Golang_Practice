package main

import "fmt"

func main() {
	fmt.Println("Hello Go users, we are going to learn Arrays in Go")

	//First way
	var price [4]int
	fmt.Println(price)
	price[0] = 12
	price[2] = 124
	price[3] = 123
	fmt.Println(price)

	//Second Way

	names := [4]string{"Sachin Tendulkar", "Rakesh Roshan", "Devdas"}
	fmt.Printf("%q", names)
	fmt.Println()

	var names2 [4]string = [4]string{"SAFK", "AKDJF"}
	fmt.Println(names2)

	//Concept of copied array and modification
	originalArray := [3]int{11, 22, 33}
	copiedArray := originalArray
	copiedArray[2] = 330
	fmt.Println("Original Array is :", originalArray)
	fmt.Println("Copied Array is :", copiedArray)

	//Iterating on arrays : 1st way
	for i := 0; i < len(names); i++ {
		fmt.Printf("%s ", names[i])
	}
	fmt.Println()

	//Iterating on arrays: 2nd way
	for i, v := range names {
		fmt.Println("Index", i, "Value", v)
	}
	// Matrix representation eg of Array
	matrix := [2][2]int{{11, 22}, {33, 44}}
	fmt.Println(matrix)

	var matrix2 [2][2]int = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matrix2)

	// More deep dive with Pointers and Addresses on Arrays
	originalArray2 := [2]int{11, 22}
	var copiedArray2 *[2]int

	copiedArray2 = &originalArray2
	copiedArray2[1] = 220
	fmt.Println(originalArray2)
	fmt.Println(copiedArray2)

	// Check whether two arrays are equal or not
	if originalArray == copiedArray {
		fmt.Println("Both originalArray and copiedArray are equal")
	} else {
		fmt.Println("originalArray and copiedArray are not equal")
	}

}

func someFunction(a, b int) (int, int) {
	return a * b, a + b
}
