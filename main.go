package main

import (
	"fmt"
)

func main() {
	color := "red"
	
	switch {
		case color == "red":
			fmt.Println("The color is red.")
		default:
			fmt.Println("The color is not red.")
	}
}