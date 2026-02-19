package main

import "fmt"

func main() {
	fmt.Println("Panic function practice in Go")
	defer fmt.Println("Defer statement 1")
	defer fmt.Println("Defer statement 2")
	panic("Panic Statement")
	defer fmt.Println("Defer statement 3")
	fmt.Println("End of the program")

}
