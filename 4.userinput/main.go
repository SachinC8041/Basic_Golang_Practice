package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var a, b int
	// fmt.Println("Please enter first Number")
	// fmt.Scan(&a)
	// fmt.Println("Please enter second number")
	// fmt.Scan(&b)

	// addition := a + b
	// fmt.Println("Addition result is :", addition)

	var name string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter username")
	name, _ = reader.ReadString('\n')

	fmt.Println("hey gophers, our username is:", name)

}
