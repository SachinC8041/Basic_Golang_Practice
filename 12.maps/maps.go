package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("Welcome to Maps in Go")
	//create a nill map
	var nilMap map[string]int
	fmt.Println(nilMap)
	if nilMap == nil {
		fmt.Println("It is nil map")
	}

	firstMap := make(map[string]int)
	firstMap["Rohit"] = 98
	firstMap["Shikhar"] = 94
	firstMap["Virat"] = 112
	firstMap["Demo"] = 0

	fmt.Println("Printing the whole map", firstMap)

	for index, value := range firstMap {
		fmt.Printf("Index %s and value %d\n", index, value)
	}

	// check the existence of the key
	ans, isPresent := firstMap["Virat"]
	fmt.Println(ans)
	fmt.Println(isPresent)

	// delete the key and value associated with it
	// delete(firstMap, "Demo")
	// fmt.Println(firstMap)

	//create map and insert multiple values at a time
	multiMap := map[string]int{
		"Game of Thrones": 9,
		"Attack on Titan": 9,
		"DeathNote":       8,
	}
	fmt.Println(multiMap)

	//finding the length
	fmt.Println("Length of multiMap is:", len(multiMap))

	//check the Equality of two maps
	if maps.Equal(multiMap, firstMap) {
		fmt.Println("Two Maps are equal")
	}

}
