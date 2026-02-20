package main

import "fmt"

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator should not be zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("We are going to learn Error Handling and Underscore operators ")
	divisionAns, err := division(10, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Division ans", divisionAns)
	}

}
