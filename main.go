package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var fullname string = "Nova Ardiansyah"
	fmt.Println(fullname)

	fullname = "First Variable"
	fmt.Println(fullname)

	company := "Nova INC"
	fmt.Println(company)

	// integer
	var ageOne int = 21
	var ageTwo = 22
	ageThree := 23
	fmt.Println(ageOne, ageTwo, ageThree)

	// bit and memory
	var numOne int8 = 127
	var numTwo int8 = -128
	fmt.Println(numOne, numTwo)

	var numTrie uint8 = 255
	fmt.Println(numTrie)

	// float
	var numFloat float32 = 3.14
	fmt.Println(numFloat)
}
