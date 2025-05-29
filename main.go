package main

import (
	"fmt"
	"golang-crash-course/calc"
)

func main() {
	fmt.Println("Hello, World!")
	sentence := testString()
	fmt.Println(sentence)

	result := calc.Add(3, 4)
	fmt.Printf("The result of adding 3 and 4 is: %d\n", result)
}