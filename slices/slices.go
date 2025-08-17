package main

import "fmt"

func findElementBinaryIterative(slice []int, element int) int {
	var left, right int = 0, len(slice) - 1
	// Perform Binary search
	for left <= right {
		mid := (left + right) / 2
		if slice[mid] == element {
			return mid
		} else if slice[mid] < element {
			left = mid + 1 // Search in the right half
		} else {
			right = mid - 1 // Search in the left half
		}
	}
	return -1
}

func findElementBinaryRecursive(slice []int, element int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if slice[mid] == element {
		return mid
	}

	if slice[mid] < element {
		return findElementBinaryRecursive(slice, element, mid+1, right) // Search in the right half
	} else {
		return findElementBinaryRecursive(slice, element, left, mid-1) // Search in the left half
	}
}

func main() {

	// Slices in go are created using the built-in make function
	var slice []int = make([]int, 5) // Create an empty slice of integers
	for i := range slice {
		slice[i] = i + 1 // Initialize the slice with values 1 to 5
	}
	fmt.Println("Initial Slice:", slice) // Print the initial slice
	const elementToFind = 3;
	fmt.Println("Finding element", elementToFind, "in the slice using binary search:", findElementBinaryIterative(slice, elementToFind)) // Find element using binary search
	fmt.Println("Finding element", elementToFind, "in the slice using recursive binary search:", findElementBinaryRecursive(slice, elementToFind, 0, len(slice)-1)) // Find element using recursive binary search
	// Print the slice
	fmt.Println("Slice:", slice)
}
