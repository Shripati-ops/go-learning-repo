package main

import "fmt"

func VariablesAndConstants() {
	// This is the simple variable definition
	var num int = 10
	fmt.Printf("num is %d \n", num)
	// Now we are going to use the short variable declaration
	num2 := 20
	fmt.Printf("num2 is %d \n", num2)
	// We can also define multiple variables at once
	var num3, num4 int = 30, 40
	fmt.Printf("num3 is %d and num4 is %d \n", num3, num4)
	// Short variable declaration for multiple variables
	num5, num6 := 50, 60
	fmt.Printf("num5 is %d and num6 is %d \n", num5, num6)
	// Constants in Go
	const pi float64 = 3.14
	fmt.Printf("The value of pi is %f \n", pi)
	// Constants can also be defined without a type
	const e = 2.71828
	fmt.Printf("The value of e is %f \n", e)
	// Constants can be defined in a block
	const (
		gravity      = 9.81
		speedOfLight = 299792458
	)
	fmt.Printf("Gravity is %f m/s^2 and speed of light is %d m/s \n", gravity, speedOfLight)
	// Constants can also be defined using iota
	const (
		red = iota
		green
		blue
	)
	fmt.Printf("Colors: red=%d, green=%d, blue=%d \n", red, green, blue)
	// iota can be used for more complex constants
	const (
		_ = iota // ignore first value
		one
		two
		three
	)
	fmt.Printf("Numbers: one=%d, two=%d, three=%d \n", one, two, three)
	// Using iota for bit shifting
}
