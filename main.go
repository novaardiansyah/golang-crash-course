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

	multiply := calc.Multiply(3, 4)
	fmt.Printf("The result of multiplying 3 and 4 is: %d\n", multiply)

	var color string = "blue"
	fmt.Printf("The color is: %s\n", color)

	age := 30
	fmt.Printf("The age is: %d\n", age)

	score := 95.5
	fmt.Printf("The score is: %.1f\n", score)
}