package main

import "fmt"

func main() {

	// Slices in go are created using the built-in make function
	var slice []int = make([]int, 5) // Create an empty slice of integers
	for i := range slice {
		slice[i] = i + 1 // Initialize the slice with values 1 to 5
	}

	// Print the slice
	fmt.Println("Slice:", slice)
}
