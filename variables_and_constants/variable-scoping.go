package main

import "fmt"

// The global scope variable
var num int = 10

func VariableScoping() {
	// The function scope variable
	var num2 int = 20
	fmt.Printf("num is %d \n", num)
	fmt.Printf("num2 is %d \n", num2)

	// This is a block scope variable
	if true {
		var num3 int = 30
		fmt.Printf("num3 is %d \n", num3)
	}

	// Example of variable scoping
	x := 10
	fmt.Println("x outside block:", x)
	{
		y := 20
		fmt.Println("x inside block:", x)
		fmt.Println("y inside block:", y)
	}
	// fmt.Println("y outside block:", y) // This would cause an error
}
