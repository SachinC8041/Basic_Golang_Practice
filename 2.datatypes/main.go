package main

import "fmt"

func main() {
	var name string = "Sachin Tendulkar"
	var age int = 12
	var temperature float64 = 23.23
	var isApplicable = false
	const taxApplicablePercentage int = 12
	// taxApplicablePercentage = 23; // you cant reassign value if the varible is constant ie. const

	//another way to declare the variables
	count := 12
	fmt.Println(count)

	fmt.Println("Name of the person is:", name)
	fmt.Println("Age of the person is:", age)
	fmt.Println("Temperature of the city is :", temperature, "And current tax rate is:", taxApplicablePercentage)
	fmt.Println("the is wrong so :", isApplicable)

	//Println and Printf function demo

	fmt.Println("Name =", name, "Age =", age, "Temperature =", temperature, "isApplicable =", isApplicable, "TaxRate =", taxApplicablePercentage)
	fmt.Printf("Name is %s and Age is %d and temperature is %f and taxrate is %d and isApplicable is %t\n", name, age, temperature, taxApplicablePercentage, isApplicable)

}
