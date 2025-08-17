package main

import "fmt"

// The main function to demonstrate slice usage

func checkKeyExists(m map[string]string, key string) bool {
	_, ok := m[key]
	return ok
}

func main() {
	m := map[string]string{
		"name": "John",
		"age":  "30",
	} // Add another key-value pair
	keyToCheck := "name"
	fmt.Printf("Key '%s' exists: %t", keyToCheck, checkKeyExists(m, keyToCheck)) // Check if key exists

	fmt.Println("\nInitial Map:", m) // Print the initial map
}
