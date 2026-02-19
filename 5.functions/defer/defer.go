package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer function practice")

	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")
	defer fmt.Println("Statement 3")

	fmt.Println("Main statement")

	var num int
	num = 12
	defer fmt.Println(num)
	num++
	fmt.Println(num)

	defer func() {
		fmt.Println("Defer with function")
	}()
	fmt.Println("Defer with normal function statement")
}
