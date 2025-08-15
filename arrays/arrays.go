package main

import "fmt"

// The main function to demonstrate array usage
func deleteElement(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		return arr // Return the original array if index is out of bounds
	}
	println(arr[:index], "Array before said index")
	arr = append(arr[:index], arr[index+1:]...) // Remove the element at the specified index
	return arr
}

func addElement(arr []int, element int) []int {
	arr = append(arr, element) // Add the new element to the end of the array
	return arr
}

func findElementLinear(arr []int, element int) int {
	for i, v := range arr {
		if v == element {
			return i // Return the index of the element if found
		}
	}
	return -1 // Return -1 if the element is not found
}

func findElementBinaryIterative(arr []int, element int) int {
	var left, right int = 0, len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == element {
			return mid // Return the index if the element is found
		} else if arr[mid] < element {
			left = mid + 1 // Search in the right half
		} else {
			right = mid - 1 // Search in the left half
		}
	}
	return -1 // Return -1 if the element is not found
}

func findElementBinaryRecursive(arr []int, element int, left int, right int) int {
	if left > right {
		return -1 // Base case: element not found
	}
	mid := (left + right) / 2
	if arr[mid] == element {
		return mid
	}

	if arr[mid] < element {
		return findElementBinaryRecursive(arr, element, mid+1, right) // Search in the right half
	} else {
		return findElementBinaryRecursive(arr, element, left, mid-1) // Search in the left half
	}

}

func main() {
	var arr []int = []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		println("Index:", i, "Value:", v)
	}

	arr = addElement(arr, 6)                  // Add a new element 6 to the array
	fmt.Println("Array after addition:", arr) // Print the array after addition

	fmt.Println("Finding element 3 in the array:", findElementLinear(arr, 3))   // Find element 3
	fmt.Println("Finding element 10 in the array:", findElementLinear(arr, 10)) // Find element 10 (not present)

	fmt.Println("Finding element 4 in the array using binary search:", findElementBinaryIterative(arr, 4))   // Find element 4 using binary search
	fmt.Println("Finding element 10 in the array using binary search:", findElementBinaryIterative(arr, 10)) // Find element 10 (not present)

	fmt.Println("Finding element 5 in the array using binary search (recursive):", findElementBinaryRecursive(arr, 5, 0, len(arr)-1)) // Find element 5 using binary search recursively

	// Demonstrating deletion
	arr = deleteElement(arr, 0)               // Delete the first element
	fmt.Println("Array after deletion:", arr) // Print the array after deletion
}
