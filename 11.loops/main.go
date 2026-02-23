package main

import "fmt"

func main() {
	fmt.Println("We are going to learn Loops in GO")
	var inputNumber int
	fmt.Scan(&inputNumber)
	//First way
	for i := 1; i <= 10; i++ {
		ans := i * inputNumber
		fmt.Printf("%d ", ans)
	}
	fmt.Println()

	//Second Way : infinite loop
	counter := 1
	for {
		fmt.Println("Counting Values")
		if counter == 4 {
			break
		}
		counter++
	}

	//Third Way: Indexes and Values using Range
	names := []string{"Sachin Tendulkar", "Virendra Sehwag", "Gautam Gambhir", "Virat Kohli", "Yuvraj Singh", "Suresh Raina"}
	for i, v := range names {
		fmt.Printf("Index is %d and Value is %s\n", i+1, v)
	}

	//Practice loop
	var num int
	fmt.Println("Please enter a valid number")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		ans := i * num

		if ans%2 == 0 {
			continue
		}

		if ans%9 == 0 && ans%8 == 0 {
			break
		}
		fmt.Printf("%d ", ans)

	}
	fmt.Println()
}
