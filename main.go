package main

import (
	"fmt"
)

func main() {
	age := 30

	if age < 18 {
		fmt.Println("You are a minor.")
	} else {
		fmt.Println("You are an adult.")
	}
}