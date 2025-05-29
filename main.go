package main

import (
	"fmt"
)

func main() {
	// ! Simple for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// ! Using a range loop
	for i := range 5 {
		fmt.Println("Range iteration:", i)
	}
}