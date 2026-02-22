package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("We are going to learn Slices in Go")

	names := []string{"Sachin", "Rahul", "Gautam", "Virat"}
	fmt.Println(names)

	names = append(names, "Virendra")
	fmt.Println(len(names))

	names2 := make([]string, 5, 5)
	names2 = append(names2, "Raja")
	names2[0] = "Rammohan"
	fmt.Printf("%q\n", names2)
	fmt.Println(len(names2))
	fmt.Println(cap(names2))

	//Equality Check in Slices
	if slices.Equal(names, names2) {
		fmt.Println("Two slices are equal")
	}

}
